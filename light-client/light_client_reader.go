package lightclient

import (
	"context"
	"fmt"
	"math/big"

	"github.com/EspressoSystems/espresso-sequencer-go/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

type LightClientReader struct {
	LightClient Lightclient
	L1Client    bind.ContractBackend
}

func NewLightClientReader(lightClientAddr common.Address, l1client bind.ContractBackend) (*LightClientReader, error) {
	lightclient, err := NewLightclient(lightClientAddr, l1client)
	if err != nil {
		return nil, err
	}

	return &LightClientReader{
		LightClient: *lightclient,
		L1Client:    l1client,
	}, nil

}

// Returns the L1 block number where the light client has validated a particular
// hotshot block number
func (l *LightClientReader) ValidatedHeight() (validatedHeight uint64, l1Height uint64, err error) {
	header, err := l.L1Client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return 0, 0, err
	}
	state, err := l.LightClient.GetFinalizedState(&bind.CallOpts{BlockNumber: header.Number})
	fmt.Printf("%+v\n", state)
	// state, err := l.LightClient.GetFinalizedState(&bind.CallOpts{})
	if err != nil {
		return 0, 0, err
	}
	return state.BlockHeight, header.Number.Uint64(), nil
}

// Fetch the merkle root at a given L1 checkpoint
func (l *LightClientReader) FetchMerkleRootAtL1Block(L1BlockHeight uint64) (types.BlockMerkleRoot, error) {
	state, err := l.LightClient.GetFinalizedState(&bind.CallOpts{BlockNumber: new(big.Int).SetUint64(L1BlockHeight)})
	if err != nil {
		return types.Commitment{}, err
	}
	root, err := types.CommitmentFromUint256(types.NewU256().SetBigInt(state.BlockCommRoot))
	if err != nil {
		return types.Commitment{}, err
	}
	return root, nil
}
