package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/EspressoSystems/espresso-sequencer-go/types"

	"github.com/ethereum/go-ethereum/log"
)

type Client struct {
	baseUrl string
	client  *http.Client
	log     log.Logger
}

func NewClient(log log.Logger, url string) *Client {
	if !strings.HasSuffix(url, "/") {
		url += "/"
	}
	return &Client{
		baseUrl: url,
		client:  http.DefaultClient,
		log:     log,
	}
}

func (c *Client) FetchLatestBlockHeight(ctx context.Context) (uint64, error) {
	var res uint64
	if err := c.get(ctx, &res, "status/latest_block_height"); err != nil {
		return 0, err
	}
	return res, nil
}

func (c *Client) FetchHeadersForWindow(ctx context.Context, start uint64, end uint64) (WindowStart, error) {
	var res WindowStart
	if err := c.get(ctx, &res, "availability/headers/window/%d/%d", start, end); err != nil {
		return WindowStart{}, err
	}
	return res, nil
}

func (c *Client) FetchHeaderByHeight(ctx context.Context, blockHeight uint64) (types.Header, error) {
	var res types.Header
	if err := c.get(ctx, &res, "availability/header/%d", blockHeight); err != nil {
		return types.Header{}, err
	}
	return res, nil
}

func (c *Client) FetchRemainingHeadersForWindow(ctx context.Context, from uint64, end uint64) (WindowMore, error) {
	var res WindowMore
	if err := c.get(ctx, &res, "availability/headers/window/from/%d/%d", from, end); err != nil {
		return WindowMore{}, err
	}
	return res, nil
}

func (c *Client) FetchTransactionsInBlock(ctx context.Context, header *types.Header, namespace uint64) (TransactionsInBlock, error) {
	var res NamespaceResponse
	if err := c.get(ctx, &res, "availability/block/%d/namespace/%d", header.Height, namespace); err != nil {
		return TransactionsInBlock{}, err
	}
	return res.Validate(header, namespace)
}

func (c *Client) SubmitTransaction(ctx context.Context, tx types.Transaction) error {
	marshalled, err := json.Marshal(tx)
	if err != nil {
		return err
	}

	request, err := http.NewRequestWithContext(ctx, "POST", c.baseUrl+"submit/submit", bytes.NewBuffer(marshalled))
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", "application/json")
	response, err := c.client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		return fmt.Errorf("received unexpected status code: %v", response.StatusCode)
	}
	return nil
}

type NamespaceResponse struct {
	Proof        *json.RawMessage     `json:"proof"`
	Transactions *[]types.Transaction `json:"transactions"`
}

// Validate a NamespaceResponse and extract the transactions.
// NMT proof validation is currently stubbed out.
func (res *NamespaceResponse) Validate(header *types.Header, namespace uint64) (TransactionsInBlock, error) {
	if res.Proof == nil {
		return TransactionsInBlock{}, fmt.Errorf("field proof of type NamespaceResponse is required")
	}
	if res.Transactions == nil {
		return TransactionsInBlock{}, fmt.Errorf("field transactions of type NamespaceResponse is required")
	}

	// Check that these transactions are only and all of the transactions from `namespace` in the
	// block with `header`.
	// TODO this is a hack. We should use the proof from the response (`proof := NmtProof{}`).
	// However, due to a simplification in the Espresso NMT implementation, where left and right
	// boundary transactions not belonging to this namespace are included in the proof in their
	// entirety, this proof can be quite large, even if this rollup has no large transactions in its
	// own namespace. In production, we have run into issues where huge transactions from other
	// rollups cause this proof to be so large, that the resulting PayloadAttributes exceeds the
	// maximum size allowed for an HTTP request by OP geth. Since NMT proof validation is currently
	// mocked anyways, we can subvert this issue in the short term without making the rollup any
	// less secure than it already is simply by using an empty proof.
	proof := types.NmtProof{}
	if err := proof.Validate(header.TransactionsRoot, *res.Transactions); err != nil {
		return TransactionsInBlock{}, err
	}

	// Extract the transactions.
	var txs []types.Bytes
	for i, tx := range *res.Transactions {
		if tx.Vm != namespace {
			return TransactionsInBlock{}, fmt.Errorf("transaction %d has wrong namespace (%d, expected %d)", i, tx.Vm, namespace)
		}
		txs = append(txs, tx.Payload)
	}

	return TransactionsInBlock{
		Transactions: txs,
		Proof:        proof,
	}, nil
}

func (c *Client) get(ctx context.Context, out any, format string, args ...any) error {
	url := c.baseUrl + fmt.Sprintf(format, args...)

	c.log.Debug("get", "url", url)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		c.log.Error("failed to build request", "err", err, "url", url)
		return err
	}
	res, err := c.client.Do(req)
	if err != nil {
		c.log.Error("error in request", "err", err, "url", url)
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		// Try to get the response body to include in the error message, as it may have useful
		// information about why the request failed. If this call fails, the response will be `nil`,
		// which is fine to include in the log, so we can ignore errors.
		body, _ := io.ReadAll(res.Body)
		c.log.Error("request failed", "err", err, "url", url, "status", res.StatusCode, "response", string(body))
		return fmt.Errorf("request failed with status %d", res.StatusCode)
	}

	// Read the response body into memory before we unmarshal it, rather than passing the io.Reader
	// to the json decoder, so that we still have the body and can inspect it if unmarshalling
	// failed.
	body, err := io.ReadAll(res.Body)
	if err != nil {
		c.log.Error("failed to read response body", "err", err, "url", url)
		return err
	}
	if err := json.Unmarshal(body, out); err != nil {
		c.log.Error("failed to parse body as json", "err", err, "url", url, "response", string(body))
		return err
	}
	c.log.Debug("request completed successfully", "url", url, "res", res, "body", string(body), "out", out)
	return nil
}
