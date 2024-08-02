package client

import (
	"context"

	"github.com/EspressoSystems/espresso-sequencer-go/types"
)

// Interface to the Espresso Sequencer submit API
type SubmitAPI interface {
	// Submit a transaction to the espresso sequencer.
	SubmitTransaction(ctx context.Context, tx types.Transaction) (*types.TaggedBase64, error)
}
