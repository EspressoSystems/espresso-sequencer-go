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
)

type Client struct {
	baseUrl string
	client  *http.Client
}

func NewClient(url string) *Client {
	if !strings.HasSuffix(url, "/") {
		url += "/"
	}
	return &Client{
		baseUrl: url,
		client:  http.DefaultClient,
	}
}

func (c *Client) FetchVidCommonByHeight(ctx context.Context, blockHeight uint64) (types.VidCommon, error) {
	var res types.VidCommonQueryData
	if err := c.get(ctx, &res, "availability/vid/common/%d", blockHeight); err != nil {
		return types.VidCommon{}, err
	}
	return res.Common, nil
}

func (c *Client) FetchLatestBlockHeight(ctx context.Context) (uint64, error) {
	var res uint64
	if err := c.get(ctx, &res, "status/block-height"); err != nil {
		return 0, err
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

func (c *Client) FetchHeadersByRange(ctx context.Context, from uint64, until uint64) ([]types.Header, error) {
	var res []types.Header
	if err := c.get(ctx, &res, "availability/header/%d/%d", from, until); err != nil {
		return []types.Header{}, err
	}
	return res, nil
}

// Fetches a block merkle proof at the snapshot rootHeight for the leaf at the provided HotShot height
func (c *Client) FetchBlockMerkleProof(ctx context.Context, rootHeight uint64, hotshotHeight uint64) (types.HotShotBlockMerkleProof, error) {
	var res types.HotShotBlockMerkleProof
	if err := c.get(ctx, &res, "block-state/%d/%d", rootHeight, hotshotHeight); err != nil {
		return types.HotShotBlockMerkleProof{}, err
	}
	return res, nil
}

func (c *Client) FetchTransactionsInBlock(ctx context.Context, blockHeight uint64, namespace uint64) (TransactionsInBlock, error) {
	var res NamespaceResponse
	if err := c.get(ctx, &res, "availability/block/%d/namespace/%d", blockHeight, namespace); err != nil {
		return TransactionsInBlock{}, err
	}

	if res.Transactions == nil {
		return TransactionsInBlock{}, fmt.Errorf("field transactions of type NamespaceResponse is required")
	}

	// Extract the transactions.
	var txs []types.Bytes
	for i, tx := range *res.Transactions {
		if tx.Namespace != namespace {
			return TransactionsInBlock{}, fmt.Errorf("transaction %d has wrong namespace (%d, expected %d)", i, tx.Namespace, namespace)
		}
		txs = append(txs, tx.Payload)
	}

	if len(txs) > 0 && res.Proof == nil {
		return TransactionsInBlock{}, fmt.Errorf("field proof of type NamespaceResponse is required")
	}

	if res.Proof == nil {
		return TransactionsInBlock{}, nil
	}

	vidCommon, err := c.FetchVidCommonByHeight(ctx, blockHeight)
	if err != nil {
		return TransactionsInBlock{}, err
	}

	return TransactionsInBlock{
		Transactions: txs,
		Proof:        *res.Proof,
		VidCommon:    vidCommon,
	}, nil

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

func (c *Client) get(ctx context.Context, out any, format string, args ...any) error {
	url := c.baseUrl + fmt.Sprintf(format, args...)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	res, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		// Try to get the response body to include in the error message, as it may have useful
		// information about why the request failed. If this call fails, the response will be `nil`,
		// which is fine to include in the log, so we can ignore errors.
		body, _ := io.ReadAll(res.Body)
		return fmt.Errorf("request failed with status %d and body %s", res.StatusCode, string(body))
	}

	// Read the response body into memory before we unmarshal it, rather than passing the io.Reader
	// to the json decoder, so that we still have the body and can inspect it if unmarshalling
	// failed.
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(body, out); err != nil {
		return err
	}
	return nil
}
