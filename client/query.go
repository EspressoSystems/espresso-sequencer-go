package client

import (
	"context"
	"encoding/json"
	"fmt"

	types "github.com/EspressoSystems/espresso-sequencer-go/types"
	common "github.com/EspressoSystems/espresso-sequencer-go/types/common"
)

// Interface to the Espresso Sequencer query service.
type QueryService interface {
	// Get the latest block number
	FetchLatestBlockHeight(ctx context.Context) (uint64, error)
	// Get the header for block number `height`.
	FetchHeaderByHeight(ctx context.Context, height uint64) (types.HeaderImpl, error)
	// Get the headers starting from the given :from up until the given :until
	FetchHeadersByRange(ctx context.Context, from uint64, until uint64) ([]types.HeaderImpl, error)
	// Get the transactions belonging to the given namespace at the block height,
	// along with a proof that these are all such transactions.
	FetchTransactionsInBlock(ctx context.Context, blockHeight uint64, namespace uint64) (TransactionsInBlock, error)
	// Get the transaction by its hash.
	FetchTransactionByHash(ctx context.Context, hash *common.TaggedBase64) (common.TransactionQueryData, error)
}

// Response to `FetchTransactionsInBlock`
type TransactionsInBlock struct {
	// The transactions.
	Transactions []common.Bytes `json:"transactions"`
	// A proof that these are all the transactions in the block with the requested namespace.
	Proof     common.NamespaceProof `json:"proof"`
	VidCommon common.VidCommon      `json:"vidCommon"`
}

func (t *TransactionsInBlock) UnmarshalJSON(b []byte) error {
	// Parse using pointers so we can distinguish between missing and default fields.
	type Dec struct {
		Transactions *[]common.Bytes        `json:"transactions"`
		Proof        *common.NamespaceProof `json:"proof"`
	}

	var dec Dec
	if err := json.Unmarshal(b, &dec); err != nil {
		return err
	}

	if dec.Transactions == nil {
		return fmt.Errorf("Field transactions of type TransactionsInBlock is required")
	}
	t.Transactions = *dec.Transactions

	if dec.Proof == nil {
		return fmt.Errorf("Field proof of type TransactionsInBlock is required")
	}
	t.Proof = *dec.Proof

	return nil
}
