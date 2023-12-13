package client

import (
	"context"
	"github.com/EspressoSystems/espresso-sequencer-go/types"
)

// Interface to the Espresso Sequencer submit API
type SubmitAPI interface {
	// Submit a transaction to the espresso sequencer with namespace `namespace`.
	SubmitTransaction(ctx context.Context, namespace uint64, tx types.Transaction) error
}
