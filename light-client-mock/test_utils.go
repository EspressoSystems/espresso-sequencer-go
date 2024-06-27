package lightclientmock

import (
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func FreezeFinalizedState(t *testing.T, client bind.ContractBackend, address common.Address, txOpts *bind.TransactOpts) error {
	mockLightClient, err := NewLightclientmock(address, client)

	if err != nil {
		return err
	}
	state, err := mockLightClient.GetFinalizedState(nil)

	if err != nil {
		return err
	}

	state.BlockHeight += 1

	mockLightClient.SetFinalizedState(txOpts, state)
	return nil
}
