// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lightclient

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// BN254G1Point is an auto generated low-level Go binding around an user-defined struct.
type BN254G1Point struct {
	X *big.Int
	Y *big.Int
}

// IPlonkVerifierPlonkProof is an auto generated low-level Go binding around an user-defined struct.
type IPlonkVerifierPlonkProof struct {
	Wire0                 BN254G1Point
	Wire1                 BN254G1Point
	Wire2                 BN254G1Point
	Wire3                 BN254G1Point
	Wire4                 BN254G1Point
	ProdPerm              BN254G1Point
	Split0                BN254G1Point
	Split1                BN254G1Point
	Split2                BN254G1Point
	Split3                BN254G1Point
	Split4                BN254G1Point
	Zeta                  BN254G1Point
	ZetaOmega             BN254G1Point
	WireEval0             *big.Int
	WireEval1             *big.Int
	WireEval2             *big.Int
	WireEval3             *big.Int
	WireEval4             *big.Int
	SigmaEval0            *big.Int
	SigmaEval1            *big.Int
	SigmaEval2            *big.Int
	SigmaEval3            *big.Int
	ProdPermZetaOmegaEval *big.Int
}

// LightClientLightClientState is an auto generated low-level Go binding around an user-defined struct.
type LightClientLightClientState struct {
	ViewNum                  uint64
	BlockHeight              uint64
	BlockCommRoot            *big.Int
	FeeLedgerComm            *big.Int
	StakeTableBlsKeyComm     *big.Int
	StakeTableSchnorrKeyComm *big.Int
	StakeTableAmountComm     *big.Int
	Threshold                *big.Int
}

// LightclientMetaData contains all meta data concerning the Lightclient contract.
var LightclientMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"genesis\",\"type\":\"tuple\",\"internalType\":\"structLightClient.LightClientState\",\"components\":[{\"name\":\"viewNum\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"blockHeight\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"blockCommRoot\",\"type\":\"uint256\",\"internalType\":\"BN254.ScalarField\"},{\"name\":\"feeLedgerComm\",\"type\":\"uint256\",\"internalType\":\"BN254.ScalarField\"},{\"name\":\"stakeTableBlsKeyComm\",\"type\":\"uint256\",\"internalType\":\"BN254.ScalarField\"},{\"name\":\"stakeTableSchnorrKeyComm\",\"type\":\"uint256\",\"internalType\":\"BN254.ScalarField\"},{\"name\":\"stakeTableAmountComm\",\"type\":\"uint256\",\"internalType\":\"BN254.ScalarField\"},{\"name\":\"threshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"numBlockPerEpoch\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"BLOCKS_PER_EPOCH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"computeStakeTableComm\",\"inputs\":[{\"name\":\"state\",\"type\":\"tuple\",\"internalType\":\"structLightClient.LightClientState\",\"components\":[{\"name\":\"viewNum\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"blockHeight\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"blockCommRoot\",\"type\":\"uint256\",\"internalType\":\"BN254.ScalarField\"},{\"name\":\"feeLedgerComm\",\"type\":\"uint256\",\"internalType\":\"BN254.ScalarField\"},{\"name\":\"stakeTableBlsKeyComm\",\"type\":\"uint256\",\"internalType\":\"BN254.ScalarField\"},{\"name\":\"stakeTableSchnorrKeyComm\",\"type\":\"uint256\",\"internalType\":\"BN254.ScalarField\"},{\"name\":\"stakeTableAmountComm\",\"type\":\"uint256\",\"internalType\":\"BN254.ScalarField\"},{\"name\":\"threshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"currentEpoch\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"finalizedState\",\"inputs\":[],\"outputs\":[{\"name\":\"viewNum\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"blockHeight\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"blockCommRoot\",\"type\":\"uint256\",\"internalType\":\"BN254.ScalarField\"},{\"name\":\"feeLedgerComm\",\"type\":\"uint256\",\"internalType\":\"BN254.ScalarField\"},{\"name\":\"stakeTableBlsKeyComm\",\"type\":\"uint256\",\"internalType\":\"BN254.ScalarField\"},{\"name\":\"stakeTableSchnorrKeyComm\",\"type\":\"uint256\",\"internalType\":\"BN254.ScalarField\"},{\"name\":\"stakeTableAmountComm\",\"type\":\"uint256\",\"internalType\":\"BN254.ScalarField\"},{\"name\":\"threshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"frozenStakeTableCommitment\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"frozenThreshold\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"genesisState\",\"inputs\":[],\"outputs\":[{\"name\":\"viewNum\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"blockHeight\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"blockCommRoot\",\"type\":\"uint256\",\"internalType\":\"BN254.ScalarField\"},{\"name\":\"feeLedgerComm\",\"type\":\"uint256\",\"internalType\":\"BN254.ScalarField\"},{\"name\":\"stakeTableBlsKeyComm\",\"type\":\"uint256\",\"internalType\":\"BN254.ScalarField\"},{\"name\":\"stakeTableSchnorrKeyComm\",\"type\":\"uint256\",\"internalType\":\"BN254.ScalarField\"},{\"name\":\"stakeTableAmountComm\",\"type\":\"uint256\",\"internalType\":\"BN254.ScalarField\"},{\"name\":\"threshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"newFinalizedState\",\"inputs\":[{\"name\":\"newState\",\"type\":\"tuple\",\"internalType\":\"structLightClient.LightClientState\",\"components\":[{\"name\":\"viewNum\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"blockHeight\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"blockCommRoot\",\"type\":\"uint256\",\"internalType\":\"BN254.ScalarField\"},{\"name\":\"feeLedgerComm\",\"type\":\"uint256\",\"internalType\":\"BN254.ScalarField\"},{\"name\":\"stakeTableBlsKeyComm\",\"type\":\"uint256\",\"internalType\":\"BN254.ScalarField\"},{\"name\":\"stakeTableSchnorrKeyComm\",\"type\":\"uint256\",\"internalType\":\"BN254.ScalarField\"},{\"name\":\"stakeTableAmountComm\",\"type\":\"uint256\",\"internalType\":\"BN254.ScalarField\"},{\"name\":\"threshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"proof\",\"type\":\"tuple\",\"internalType\":\"structIPlonkVerifier.PlonkProof\",\"components\":[{\"name\":\"wire0\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"x\",\"type\":\"uint256\",\"internalType\":\"BN254.BaseField\"},{\"name\":\"y\",\"type\":\"uint256\",\"internalType\":\"BN254.BaseField\"}]},{\"name\":\"wire1\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"x\",\"type\":\"uint256\",\"internalType\":\"BN254.BaseField\"},{\"name\":\"y\",\"type\":\"uint256\",\"internalType\":\"BN254.BaseField\"}]},{\"name\":\"wire2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"x\",\"type\":\"uint256\",\"internalType\":\"BN254.BaseField\"},{\"name\":\"y\",\"type\":\"uint256\",\"internalType\":\"BN254.BaseField\"}]},{\"name\":\"wire3\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"x\",\"type\":\"uint256\",\"internalType\":\"BN254.BaseField\"},{\"name\":\"y\",\"type\":\"uint256\",\"internalType\":\"BN254.BaseField\"}]},{\"name\":\"wire4\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"x\",\"type\":\"uint256\",\"internalType\":\"BN254.BaseField\"},{\"name\":\"y\",\"type\":\"uint256\",\"internalType\":\"BN254.BaseField\"}]},{\"name\":\"prodPerm\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"x\",\"type\":\"uint256\",\"internalType\":\"BN254.BaseField\"},{\"name\":\"y\",\"type\":\"uint256\",\"internalType\":\"BN254.BaseField\"}]},{\"name\":\"split0\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"x\",\"type\":\"uint256\",\"internalType\":\"BN254.BaseField\"},{\"name\":\"y\",\"type\":\"uint256\",\"internalType\":\"BN254.BaseField\"}]},{\"name\":\"split1\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"x\",\"type\":\"uint256\",\"internalType\":\"BN254.BaseField\"},{\"name\":\"y\",\"type\":\"uint256\",\"internalType\":\"BN254.BaseField\"}]},{\"name\":\"split2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"x\",\"type\":\"uint256\",\"internalType\":\"BN254.BaseField\"},{\"name\":\"y\",\"type\":\"uint256\",\"internalType\":\"BN254.BaseField\"}]},{\"name\":\"split3\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"x\",\"type\":\"uint256\",\"internalType\":\"BN254.BaseField\"},{\"name\":\"y\",\"type\":\"uint256\",\"internalType\":\"BN254.BaseField\"}]},{\"name\":\"split4\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"x\",\"type\":\"uint256\",\"internalType\":\"BN254.BaseField\"},{\"name\":\"y\",\"type\":\"uint256\",\"internalType\":\"BN254.BaseField\"}]},{\"name\":\"zeta\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"x\",\"type\":\"uint256\",\"internalType\":\"BN254.BaseField\"},{\"name\":\"y\",\"type\":\"uint256\",\"internalType\":\"BN254.BaseField\"}]},{\"name\":\"zetaOmega\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"x\",\"type\":\"uint256\",\"internalType\":\"BN254.BaseField\"},{\"name\":\"y\",\"type\":\"uint256\",\"internalType\":\"BN254.BaseField\"}]},{\"name\":\"wireEval0\",\"type\":\"uint256\",\"internalType\":\"BN254.ScalarField\"},{\"name\":\"wireEval1\",\"type\":\"uint256\",\"internalType\":\"BN254.ScalarField\"},{\"name\":\"wireEval2\",\"type\":\"uint256\",\"internalType\":\"BN254.ScalarField\"},{\"name\":\"wireEval3\",\"type\":\"uint256\",\"internalType\":\"BN254.ScalarField\"},{\"name\":\"wireEval4\",\"type\":\"uint256\",\"internalType\":\"BN254.ScalarField\"},{\"name\":\"sigmaEval0\",\"type\":\"uint256\",\"internalType\":\"BN254.ScalarField\"},{\"name\":\"sigmaEval1\",\"type\":\"uint256\",\"internalType\":\"BN254.ScalarField\"},{\"name\":\"sigmaEval2\",\"type\":\"uint256\",\"internalType\":\"BN254.ScalarField\"},{\"name\":\"sigmaEval3\",\"type\":\"uint256\",\"internalType\":\"BN254.ScalarField\"},{\"name\":\"prodPermZetaOmegaEval\",\"type\":\"uint256\",\"internalType\":\"BN254.ScalarField\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"votingStakeTableCommitment\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"votingThreshold\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"EpochChanged\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NewState\",\"inputs\":[{\"name\":\"viewNum\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"blockHeight\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"blockCommRoot\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"BN254.ScalarField\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidArgs\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidProof\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MissingLastBlockForCurrentEpoch\",\"inputs\":[{\"name\":\"expectedBlockHeight\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"OutdatedState\",\"inputs\":[]}]",
}

// LightclientABI is the input ABI used to generate the binding from.
// Deprecated: Use LightclientMetaData.ABI instead.
var LightclientABI = LightclientMetaData.ABI

// Lightclient is an auto generated Go binding around an Ethereum contract.
type Lightclient struct {
	LightclientCaller     // Read-only binding to the contract
	LightclientTransactor // Write-only binding to the contract
	LightclientFilterer   // Log filterer for contract events
}

// LightclientCaller is an auto generated read-only Go binding around an Ethereum contract.
type LightclientCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LightclientTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LightclientTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LightclientFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LightclientFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LightclientSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LightclientSession struct {
	Contract     *Lightclient      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LightclientCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LightclientCallerSession struct {
	Contract *LightclientCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// LightclientTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LightclientTransactorSession struct {
	Contract     *LightclientTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// LightclientRaw is an auto generated low-level Go binding around an Ethereum contract.
type LightclientRaw struct {
	Contract *Lightclient // Generic contract binding to access the raw methods on
}

// LightclientCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LightclientCallerRaw struct {
	Contract *LightclientCaller // Generic read-only contract binding to access the raw methods on
}

// LightclientTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LightclientTransactorRaw struct {
	Contract *LightclientTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLightclient creates a new instance of Lightclient, bound to a specific deployed contract.
func NewLightclient(address common.Address, backend bind.ContractBackend) (*Lightclient, error) {
	contract, err := bindLightclient(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Lightclient{LightclientCaller: LightclientCaller{contract: contract}, LightclientTransactor: LightclientTransactor{contract: contract}, LightclientFilterer: LightclientFilterer{contract: contract}}, nil
}

// NewLightclientCaller creates a new read-only instance of Lightclient, bound to a specific deployed contract.
func NewLightclientCaller(address common.Address, caller bind.ContractCaller) (*LightclientCaller, error) {
	contract, err := bindLightclient(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LightclientCaller{contract: contract}, nil
}

// NewLightclientTransactor creates a new write-only instance of Lightclient, bound to a specific deployed contract.
func NewLightclientTransactor(address common.Address, transactor bind.ContractTransactor) (*LightclientTransactor, error) {
	contract, err := bindLightclient(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LightclientTransactor{contract: contract}, nil
}

// NewLightclientFilterer creates a new log filterer instance of Lightclient, bound to a specific deployed contract.
func NewLightclientFilterer(address common.Address, filterer bind.ContractFilterer) (*LightclientFilterer, error) {
	contract, err := bindLightclient(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LightclientFilterer{contract: contract}, nil
}

// bindLightclient binds a generic wrapper to an already deployed contract.
func bindLightclient(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LightclientMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lightclient *LightclientRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Lightclient.Contract.LightclientCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lightclient *LightclientRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lightclient.Contract.LightclientTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lightclient *LightclientRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lightclient.Contract.LightclientTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lightclient *LightclientCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Lightclient.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lightclient *LightclientTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lightclient.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lightclient *LightclientTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lightclient.Contract.contract.Transact(opts, method, params...)
}

// BLOCKSPEREPOCH is a free data retrieval call binding the contract method 0x1f3adc0f.
//
// Solidity: function BLOCKS_PER_EPOCH() view returns(uint32)
func (_Lightclient *LightclientCaller) BLOCKSPEREPOCH(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Lightclient.contract.Call(opts, &out, "BLOCKS_PER_EPOCH")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// BLOCKSPEREPOCH is a free data retrieval call binding the contract method 0x1f3adc0f.
//
// Solidity: function BLOCKS_PER_EPOCH() view returns(uint32)
func (_Lightclient *LightclientSession) BLOCKSPEREPOCH() (uint32, error) {
	return _Lightclient.Contract.BLOCKSPEREPOCH(&_Lightclient.CallOpts)
}

// BLOCKSPEREPOCH is a free data retrieval call binding the contract method 0x1f3adc0f.
//
// Solidity: function BLOCKS_PER_EPOCH() view returns(uint32)
func (_Lightclient *LightclientCallerSession) BLOCKSPEREPOCH() (uint32, error) {
	return _Lightclient.Contract.BLOCKSPEREPOCH(&_Lightclient.CallOpts)
}

// ComputeStakeTableComm is a free data retrieval call binding the contract method 0xaa922732.
//
// Solidity: function computeStakeTableComm((uint64,uint64,uint256,uint256,uint256,uint256,uint256,uint256) state) pure returns(bytes32)
func (_Lightclient *LightclientCaller) ComputeStakeTableComm(opts *bind.CallOpts, state LightClientLightClientState) ([32]byte, error) {
	var out []interface{}
	err := _Lightclient.contract.Call(opts, &out, "computeStakeTableComm", state)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ComputeStakeTableComm is a free data retrieval call binding the contract method 0xaa922732.
//
// Solidity: function computeStakeTableComm((uint64,uint64,uint256,uint256,uint256,uint256,uint256,uint256) state) pure returns(bytes32)
func (_Lightclient *LightclientSession) ComputeStakeTableComm(state LightClientLightClientState) ([32]byte, error) {
	return _Lightclient.Contract.ComputeStakeTableComm(&_Lightclient.CallOpts, state)
}

// ComputeStakeTableComm is a free data retrieval call binding the contract method 0xaa922732.
//
// Solidity: function computeStakeTableComm((uint64,uint64,uint256,uint256,uint256,uint256,uint256,uint256) state) pure returns(bytes32)
func (_Lightclient *LightclientCallerSession) ComputeStakeTableComm(state LightClientLightClientState) ([32]byte, error) {
	return _Lightclient.Contract.ComputeStakeTableComm(&_Lightclient.CallOpts, state)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint64)
func (_Lightclient *LightclientCaller) CurrentEpoch(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Lightclient.contract.Call(opts, &out, "currentEpoch")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint64)
func (_Lightclient *LightclientSession) CurrentEpoch() (uint64, error) {
	return _Lightclient.Contract.CurrentEpoch(&_Lightclient.CallOpts)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint64)
func (_Lightclient *LightclientCallerSession) CurrentEpoch() (uint64, error) {
	return _Lightclient.Contract.CurrentEpoch(&_Lightclient.CallOpts)
}

// FinalizedState is a free data retrieval call binding the contract method 0x9fdb54a7.
//
// Solidity: function finalizedState() view returns(uint64 viewNum, uint64 blockHeight, uint256 blockCommRoot, uint256 feeLedgerComm, uint256 stakeTableBlsKeyComm, uint256 stakeTableSchnorrKeyComm, uint256 stakeTableAmountComm, uint256 threshold)
func (_Lightclient *LightclientCaller) FinalizedState(opts *bind.CallOpts) (struct {
	ViewNum                  uint64
	BlockHeight              uint64
	BlockCommRoot            *big.Int
	FeeLedgerComm            *big.Int
	StakeTableBlsKeyComm     *big.Int
	StakeTableSchnorrKeyComm *big.Int
	StakeTableAmountComm     *big.Int
	Threshold                *big.Int
}, error) {
	var out []interface{}
	err := _Lightclient.contract.Call(opts, &out, "finalizedState")

	outstruct := new(struct {
		ViewNum                  uint64
		BlockHeight              uint64
		BlockCommRoot            *big.Int
		FeeLedgerComm            *big.Int
		StakeTableBlsKeyComm     *big.Int
		StakeTableSchnorrKeyComm *big.Int
		StakeTableAmountComm     *big.Int
		Threshold                *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ViewNum = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.BlockHeight = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.BlockCommRoot = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.FeeLedgerComm = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.StakeTableBlsKeyComm = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.StakeTableSchnorrKeyComm = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.StakeTableAmountComm = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.Threshold = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// FinalizedState is a free data retrieval call binding the contract method 0x9fdb54a7.
//
// Solidity: function finalizedState() view returns(uint64 viewNum, uint64 blockHeight, uint256 blockCommRoot, uint256 feeLedgerComm, uint256 stakeTableBlsKeyComm, uint256 stakeTableSchnorrKeyComm, uint256 stakeTableAmountComm, uint256 threshold)
func (_Lightclient *LightclientSession) FinalizedState() (struct {
	ViewNum                  uint64
	BlockHeight              uint64
	BlockCommRoot            *big.Int
	FeeLedgerComm            *big.Int
	StakeTableBlsKeyComm     *big.Int
	StakeTableSchnorrKeyComm *big.Int
	StakeTableAmountComm     *big.Int
	Threshold                *big.Int
}, error) {
	return _Lightclient.Contract.FinalizedState(&_Lightclient.CallOpts)
}

// FinalizedState is a free data retrieval call binding the contract method 0x9fdb54a7.
//
// Solidity: function finalizedState() view returns(uint64 viewNum, uint64 blockHeight, uint256 blockCommRoot, uint256 feeLedgerComm, uint256 stakeTableBlsKeyComm, uint256 stakeTableSchnorrKeyComm, uint256 stakeTableAmountComm, uint256 threshold)
func (_Lightclient *LightclientCallerSession) FinalizedState() (struct {
	ViewNum                  uint64
	BlockHeight              uint64
	BlockCommRoot            *big.Int
	FeeLedgerComm            *big.Int
	StakeTableBlsKeyComm     *big.Int
	StakeTableSchnorrKeyComm *big.Int
	StakeTableAmountComm     *big.Int
	Threshold                *big.Int
}, error) {
	return _Lightclient.Contract.FinalizedState(&_Lightclient.CallOpts)
}

// FrozenStakeTableCommitment is a free data retrieval call binding the contract method 0x382b215a.
//
// Solidity: function frozenStakeTableCommitment() view returns(bytes32)
func (_Lightclient *LightclientCaller) FrozenStakeTableCommitment(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Lightclient.contract.Call(opts, &out, "frozenStakeTableCommitment")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// FrozenStakeTableCommitment is a free data retrieval call binding the contract method 0x382b215a.
//
// Solidity: function frozenStakeTableCommitment() view returns(bytes32)
func (_Lightclient *LightclientSession) FrozenStakeTableCommitment() ([32]byte, error) {
	return _Lightclient.Contract.FrozenStakeTableCommitment(&_Lightclient.CallOpts)
}

// FrozenStakeTableCommitment is a free data retrieval call binding the contract method 0x382b215a.
//
// Solidity: function frozenStakeTableCommitment() view returns(bytes32)
func (_Lightclient *LightclientCallerSession) FrozenStakeTableCommitment() ([32]byte, error) {
	return _Lightclient.Contract.FrozenStakeTableCommitment(&_Lightclient.CallOpts)
}

// FrozenThreshold is a free data retrieval call binding the contract method 0xca6fe855.
//
// Solidity: function frozenThreshold() view returns(uint256)
func (_Lightclient *LightclientCaller) FrozenThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lightclient.contract.Call(opts, &out, "frozenThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FrozenThreshold is a free data retrieval call binding the contract method 0xca6fe855.
//
// Solidity: function frozenThreshold() view returns(uint256)
func (_Lightclient *LightclientSession) FrozenThreshold() (*big.Int, error) {
	return _Lightclient.Contract.FrozenThreshold(&_Lightclient.CallOpts)
}

// FrozenThreshold is a free data retrieval call binding the contract method 0xca6fe855.
//
// Solidity: function frozenThreshold() view returns(uint256)
func (_Lightclient *LightclientCallerSession) FrozenThreshold() (*big.Int, error) {
	return _Lightclient.Contract.FrozenThreshold(&_Lightclient.CallOpts)
}

// GenesisState is a free data retrieval call binding the contract method 0xd24d933d.
//
// Solidity: function genesisState() view returns(uint64 viewNum, uint64 blockHeight, uint256 blockCommRoot, uint256 feeLedgerComm, uint256 stakeTableBlsKeyComm, uint256 stakeTableSchnorrKeyComm, uint256 stakeTableAmountComm, uint256 threshold)
func (_Lightclient *LightclientCaller) GenesisState(opts *bind.CallOpts) (struct {
	ViewNum                  uint64
	BlockHeight              uint64
	BlockCommRoot            *big.Int
	FeeLedgerComm            *big.Int
	StakeTableBlsKeyComm     *big.Int
	StakeTableSchnorrKeyComm *big.Int
	StakeTableAmountComm     *big.Int
	Threshold                *big.Int
}, error) {
	var out []interface{}
	err := _Lightclient.contract.Call(opts, &out, "genesisState")

	outstruct := new(struct {
		ViewNum                  uint64
		BlockHeight              uint64
		BlockCommRoot            *big.Int
		FeeLedgerComm            *big.Int
		StakeTableBlsKeyComm     *big.Int
		StakeTableSchnorrKeyComm *big.Int
		StakeTableAmountComm     *big.Int
		Threshold                *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ViewNum = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.BlockHeight = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.BlockCommRoot = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.FeeLedgerComm = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.StakeTableBlsKeyComm = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.StakeTableSchnorrKeyComm = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.StakeTableAmountComm = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.Threshold = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GenesisState is a free data retrieval call binding the contract method 0xd24d933d.
//
// Solidity: function genesisState() view returns(uint64 viewNum, uint64 blockHeight, uint256 blockCommRoot, uint256 feeLedgerComm, uint256 stakeTableBlsKeyComm, uint256 stakeTableSchnorrKeyComm, uint256 stakeTableAmountComm, uint256 threshold)
func (_Lightclient *LightclientSession) GenesisState() (struct {
	ViewNum                  uint64
	BlockHeight              uint64
	BlockCommRoot            *big.Int
	FeeLedgerComm            *big.Int
	StakeTableBlsKeyComm     *big.Int
	StakeTableSchnorrKeyComm *big.Int
	StakeTableAmountComm     *big.Int
	Threshold                *big.Int
}, error) {
	return _Lightclient.Contract.GenesisState(&_Lightclient.CallOpts)
}

// GenesisState is a free data retrieval call binding the contract method 0xd24d933d.
//
// Solidity: function genesisState() view returns(uint64 viewNum, uint64 blockHeight, uint256 blockCommRoot, uint256 feeLedgerComm, uint256 stakeTableBlsKeyComm, uint256 stakeTableSchnorrKeyComm, uint256 stakeTableAmountComm, uint256 threshold)
func (_Lightclient *LightclientCallerSession) GenesisState() (struct {
	ViewNum                  uint64
	BlockHeight              uint64
	BlockCommRoot            *big.Int
	FeeLedgerComm            *big.Int
	StakeTableBlsKeyComm     *big.Int
	StakeTableSchnorrKeyComm *big.Int
	StakeTableAmountComm     *big.Int
	Threshold                *big.Int
}, error) {
	return _Lightclient.Contract.GenesisState(&_Lightclient.CallOpts)
}

// VotingStakeTableCommitment is a free data retrieval call binding the contract method 0x76b6b7cb.
//
// Solidity: function votingStakeTableCommitment() view returns(bytes32)
func (_Lightclient *LightclientCaller) VotingStakeTableCommitment(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Lightclient.contract.Call(opts, &out, "votingStakeTableCommitment")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// VotingStakeTableCommitment is a free data retrieval call binding the contract method 0x76b6b7cb.
//
// Solidity: function votingStakeTableCommitment() view returns(bytes32)
func (_Lightclient *LightclientSession) VotingStakeTableCommitment() ([32]byte, error) {
	return _Lightclient.Contract.VotingStakeTableCommitment(&_Lightclient.CallOpts)
}

// VotingStakeTableCommitment is a free data retrieval call binding the contract method 0x76b6b7cb.
//
// Solidity: function votingStakeTableCommitment() view returns(bytes32)
func (_Lightclient *LightclientCallerSession) VotingStakeTableCommitment() ([32]byte, error) {
	return _Lightclient.Contract.VotingStakeTableCommitment(&_Lightclient.CallOpts)
}

// VotingThreshold is a free data retrieval call binding the contract method 0x62827733.
//
// Solidity: function votingThreshold() view returns(uint256)
func (_Lightclient *LightclientCaller) VotingThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lightclient.contract.Call(opts, &out, "votingThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VotingThreshold is a free data retrieval call binding the contract method 0x62827733.
//
// Solidity: function votingThreshold() view returns(uint256)
func (_Lightclient *LightclientSession) VotingThreshold() (*big.Int, error) {
	return _Lightclient.Contract.VotingThreshold(&_Lightclient.CallOpts)
}

// VotingThreshold is a free data retrieval call binding the contract method 0x62827733.
//
// Solidity: function votingThreshold() view returns(uint256)
func (_Lightclient *LightclientCallerSession) VotingThreshold() (*big.Int, error) {
	return _Lightclient.Contract.VotingThreshold(&_Lightclient.CallOpts)
}

// NewFinalizedState is a paid mutator transaction binding the contract method 0x409939b7.
//
// Solidity: function newFinalizedState((uint64,uint64,uint256,uint256,uint256,uint256,uint256,uint256) newState, ((uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256),uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) proof) returns()
func (_Lightclient *LightclientTransactor) NewFinalizedState(opts *bind.TransactOpts, newState LightClientLightClientState, proof IPlonkVerifierPlonkProof) (*types.Transaction, error) {
	return _Lightclient.contract.Transact(opts, "newFinalizedState", newState, proof)
}

// NewFinalizedState is a paid mutator transaction binding the contract method 0x409939b7.
//
// Solidity: function newFinalizedState((uint64,uint64,uint256,uint256,uint256,uint256,uint256,uint256) newState, ((uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256),uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) proof) returns()
func (_Lightclient *LightclientSession) NewFinalizedState(newState LightClientLightClientState, proof IPlonkVerifierPlonkProof) (*types.Transaction, error) {
	return _Lightclient.Contract.NewFinalizedState(&_Lightclient.TransactOpts, newState, proof)
}

// NewFinalizedState is a paid mutator transaction binding the contract method 0x409939b7.
//
// Solidity: function newFinalizedState((uint64,uint64,uint256,uint256,uint256,uint256,uint256,uint256) newState, ((uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256),uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) proof) returns()
func (_Lightclient *LightclientTransactorSession) NewFinalizedState(newState LightClientLightClientState, proof IPlonkVerifierPlonkProof) (*types.Transaction, error) {
	return _Lightclient.Contract.NewFinalizedState(&_Lightclient.TransactOpts, newState, proof)
}

// LightclientEpochChangedIterator is returned from FilterEpochChanged and is used to iterate over the raw logs and unpacked data for EpochChanged events raised by the Lightclient contract.
type LightclientEpochChangedIterator struct {
	Event *LightclientEpochChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LightclientEpochChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LightclientEpochChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LightclientEpochChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LightclientEpochChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LightclientEpochChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LightclientEpochChanged represents a EpochChanged event raised by the Lightclient contract.
type LightclientEpochChanged struct {
	Arg0 uint64
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterEpochChanged is a free log retrieval operation binding the contract event 0xdb3558259e039d7e50e816b9dcce30fb114d8a9c86eca5ab14b60194d6945d3f.
//
// Solidity: event EpochChanged(uint64 arg0)
func (_Lightclient *LightclientFilterer) FilterEpochChanged(opts *bind.FilterOpts) (*LightclientEpochChangedIterator, error) {

	logs, sub, err := _Lightclient.contract.FilterLogs(opts, "EpochChanged")
	if err != nil {
		return nil, err
	}
	return &LightclientEpochChangedIterator{contract: _Lightclient.contract, event: "EpochChanged", logs: logs, sub: sub}, nil
}

// WatchEpochChanged is a free log subscription operation binding the contract event 0xdb3558259e039d7e50e816b9dcce30fb114d8a9c86eca5ab14b60194d6945d3f.
//
// Solidity: event EpochChanged(uint64 arg0)
func (_Lightclient *LightclientFilterer) WatchEpochChanged(opts *bind.WatchOpts, sink chan<- *LightclientEpochChanged) (event.Subscription, error) {

	logs, sub, err := _Lightclient.contract.WatchLogs(opts, "EpochChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LightclientEpochChanged)
				if err := _Lightclient.contract.UnpackLog(event, "EpochChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEpochChanged is a log parse operation binding the contract event 0xdb3558259e039d7e50e816b9dcce30fb114d8a9c86eca5ab14b60194d6945d3f.
//
// Solidity: event EpochChanged(uint64 arg0)
func (_Lightclient *LightclientFilterer) ParseEpochChanged(log types.Log) (*LightclientEpochChanged, error) {
	event := new(LightclientEpochChanged)
	if err := _Lightclient.contract.UnpackLog(event, "EpochChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LightclientNewStateIterator is returned from FilterNewState and is used to iterate over the raw logs and unpacked data for NewState events raised by the Lightclient contract.
type LightclientNewStateIterator struct {
	Event *LightclientNewState // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LightclientNewStateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LightclientNewState)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LightclientNewState)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LightclientNewStateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LightclientNewStateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LightclientNewState represents a NewState event raised by the Lightclient contract.
type LightclientNewState struct {
	ViewNum       uint64
	BlockHeight   uint64
	BlockCommRoot *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterNewState is a free log retrieval operation binding the contract event 0xa04a773924505a418564363725f56832f5772e6b8d0dbd6efce724dfe803dae6.
//
// Solidity: event NewState(uint64 indexed viewNum, uint64 indexed blockHeight, uint256 blockCommRoot)
func (_Lightclient *LightclientFilterer) FilterNewState(opts *bind.FilterOpts, viewNum []uint64, blockHeight []uint64) (*LightclientNewStateIterator, error) {

	var viewNumRule []interface{}
	for _, viewNumItem := range viewNum {
		viewNumRule = append(viewNumRule, viewNumItem)
	}
	var blockHeightRule []interface{}
	for _, blockHeightItem := range blockHeight {
		blockHeightRule = append(blockHeightRule, blockHeightItem)
	}

	logs, sub, err := _Lightclient.contract.FilterLogs(opts, "NewState", viewNumRule, blockHeightRule)
	if err != nil {
		return nil, err
	}
	return &LightclientNewStateIterator{contract: _Lightclient.contract, event: "NewState", logs: logs, sub: sub}, nil
}

// WatchNewState is a free log subscription operation binding the contract event 0xa04a773924505a418564363725f56832f5772e6b8d0dbd6efce724dfe803dae6.
//
// Solidity: event NewState(uint64 indexed viewNum, uint64 indexed blockHeight, uint256 blockCommRoot)
func (_Lightclient *LightclientFilterer) WatchNewState(opts *bind.WatchOpts, sink chan<- *LightclientNewState, viewNum []uint64, blockHeight []uint64) (event.Subscription, error) {

	var viewNumRule []interface{}
	for _, viewNumItem := range viewNum {
		viewNumRule = append(viewNumRule, viewNumItem)
	}
	var blockHeightRule []interface{}
	for _, blockHeightItem := range blockHeight {
		blockHeightRule = append(blockHeightRule, blockHeightItem)
	}

	logs, sub, err := _Lightclient.contract.WatchLogs(opts, "NewState", viewNumRule, blockHeightRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LightclientNewState)
				if err := _Lightclient.contract.UnpackLog(event, "NewState", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewState is a log parse operation binding the contract event 0xa04a773924505a418564363725f56832f5772e6b8d0dbd6efce724dfe803dae6.
//
// Solidity: event NewState(uint64 indexed viewNum, uint64 indexed blockHeight, uint256 blockCommRoot)
func (_Lightclient *LightclientFilterer) ParseNewState(log types.Log) (*LightclientNewState, error) {
	event := new(LightclientNewState)
	if err := _Lightclient.contract.UnpackLog(event, "NewState", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
