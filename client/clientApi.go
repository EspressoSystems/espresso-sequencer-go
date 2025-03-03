package client

import(
	"context"
	"encoding/json"

	types "github.com/EspressoSystems/espresso-sequencer-go/types"
	common "github.com/EspressoSystems/espresso-sequencer-go/types/common"
)

type EspressoClient interface{ 
  // Get the latest block number
	FetchLatestBlockHeight(ctx context.Context) (uint64, error)
	// Get the header for block number `height`.
	FetchHeaderByHeight(ctx context.Context, height uint64) (types.HeaderImpl, error)
	// Get the raw header for block number `height`.
	FetchRawHeaderByHeight(ctx context.Context, height uint64) (json.RawMessage, error)
	// Get the headers starting from the given :from up until the given :until
	FetchHeadersByRange(ctx context.Context, from uint64, until uint64) ([]types.HeaderImpl, error)
	// Get the transactions belonging to the given namespace at the block height,
	// along with a proof that these are all such transactions.
	FetchTransactionsInBlock(ctx context.Context, blockHeight uint64, namespace uint64) (TransactionsInBlock, error)
	// Get the transaction by its hash.
	FetchTransactionByHash(ctx context.Context, hash *types.TaggedBase64) (types.TransactionQueryData, error)
	// Get the VidCommon for the given block height.
	FetchVidCommonByHeight(ctx context.Context, blockHeight uint64) (types.VidCommon, error)
	// Submit a transaction to the espresso sequencer.
	SubmitTransaction(ctx context.Context, tx common.Transaction) (*common.TaggedBase64, error)
}
