package nmt

import (
	"github.com/EspressoSystems/espresso-sequencer-go/types"
)

// This function mocks batch transaction validation against a set of HotShot NMT roots.
// It pretends to verify that the set of transactions (txns) in a batch correspond to a set of n NMT proofs
// (p1, ... pn) against trusted NMT roots r1,...rn.
//
// In other words, the function validates that txns = {...p1.txns, ..., ...pn.txns}, that all the transactions
// come from the given namespace, and that p1, ..., pn are all valid NMT proofs with respect to r1, ..., rn.
func ValidateBatchTransactions(namespace uint64, roots []*types.NmtRoot, proofs []*types.NmtProof, transactions []types.Bytes) error {
	return nil
}
