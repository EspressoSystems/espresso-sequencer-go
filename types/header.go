package types

import (
	"encoding/json"
	"fmt"

	tagged_base64 "github.com/EspressoSystems/espresso-sequencer-go/tagged-base64"
	common_types "github.com/EspressoSystems/espresso-sequencer-go/types/common"
	"github.com/EspressoSystems/espresso-sequencer-go/types/v0/v0_1"
	v01 "github.com/EspressoSystems/espresso-sequencer-go/types/v0/v0_1"
	v02 "github.com/EspressoSystems/espresso-sequencer-go/types/v0/v0_2"
	v03 "github.com/EspressoSystems/espresso-sequencer-go/types/v0/v0_3"
	common "github.com/ethereum/go-ethereum/common"
)

// Republic
type EitherChainConfig0_1 = v01.EitherChainConfig
type ResolvableChainConfig0_1 = v01.ResolvableChainConfig
type ChainConfig0_1 = v01.ChainConfig

type EitherChainConfig0_3 = v03.EitherChainConfig
type ResolvableChainConfig0_3 = v03.ResolvableChainConfig
type ChainConfig0_3 = v03.ChainConfig

type Header0_1 = v01.Header
type Header0_2 = v02.Header
type Header0_3 = v03.Header

type Bytes = common_types.Bytes

type BlockMerkleSnapshot = common_types.BlockMerkleSnapshot
type Commitment = common_types.Commitment
type FeeInfo = common_types.FeeInfo
type HotShotBlockMerkleProof = common_types.HotShotBlockMerkleProof
type L1BlockInfo = common_types.L1BlockInfo
type NamespaceProof = common_types.NamespaceProof
type NsTable = common_types.NsTable
type Signature = common_types.Signature
type TaggedBase64 = common_types.TaggedBase64
type Transaction = common_types.Transaction
type Version = common_types.Version
type VidCommon = common_types.VidCommon

type TransactionQueryData = common_types.TransactionQueryData
type VidCommonQueryData = common_types.VidCommonQueryData

type U256 = common_types.U256
type U256Decimal = common_types.U256Decimal

var NewU256 = common_types.NewU256

type HeaderInterface interface {
	Commit() common_types.Commitment
	Version() common_types.Version
	GetBlockHeight() uint64
	GetL1Head() uint64
	GetTimestamp() uint64
	GetPayloadCommitment() *common_types.TaggedBase64
	GetBuilderCommitment() *common_types.TaggedBase64
	GetNsTable() *common_types.NsTable
	GetBlockMerkleTreeRoot() *common_types.TaggedBase64
	GetFeeMerkleTreeRoot() *common_types.TaggedBase64
}

type HeaderImpl struct {
	Header HeaderInterface
}

func (h *HeaderImpl) UnmarshalJSON(b []byte) error {
	header, err := parseHeader(b)
	if err != nil {
		return err
	}

	h.Header = header
	return nil
}

func (h HeaderImpl) MarshalJSON() ([]byte, error) {
	version := h.Header.Version()
	if version.Major == 0 && version.Minor == 1 {
		return json.Marshal(h.Header)
	}

	var rawHeader RawHeader
	rawHeader.Version = version

	bytes, err := json.Marshal(h.Header)
	if err != nil {
		return nil, err
	}

	rawHeader.Fields = bytes

	return json.Marshal(rawHeader)
}

type RawHeader struct {
	Version common_types.Version `json:"version"`
	Fields  json.RawMessage      `json:"fields"`
}

func (r *RawHeader) UnmarshalJSON(b []byte) error {
	type Dec struct {
		Version *common_types.Version `json:"version"`
		Fields  *json.RawMessage      `json:"fields"`
	}

	var dec Dec
	if err := json.Unmarshal(b, &dec); err != nil {
		return err
	}

	if dec.Version == nil {
		return fmt.Errorf("version is required")
	}

	if dec.Fields == nil {
		return fmt.Errorf("fields is required")
	}

	r.Version = *dec.Version
	r.Fields = *dec.Fields
	return nil
}

func parseHeader(data []byte) (HeaderInterface, error) {
	var rawHeader RawHeader
	if err := json.Unmarshal(data, &rawHeader); err != nil {
		var header v01.Header
		// Failed to parse the version header. Legacy header
		if err = json.Unmarshal(data, &header); err != nil {
			return nil, err
		}

		return &header, nil
	}

	version := rawHeader.Version
	if version.Major == 0 && version.Minor == 2 {
		var header v02.Header
		if err := json.Unmarshal(rawHeader.Fields, &header); err != nil {
			return nil, err
		}
		return &header, nil
	}

	if version.Major == 0 && version.Minor == 3 {
		var header v03.Header
		if err := json.Unmarshal(rawHeader.Fields, &header); err != nil {
			return nil, err
		}
		return &header, nil
	}

	return nil, fmt.Errorf("version error: %v", version)
}

// For some intricate reasons, rollups need to build a dummy header as a placeholder. However, an empty Header can be serialized
// but can not be deserialized because of the checks. Thus we provide this dummy header as a workaround.
func GetDummyHeader() v0_1.Header {
	var payloadCommitment, _ = tagged_base64.Parse("HASH~1yS-KEtL3oDZDBJdsW51Pd7zywIiHesBZsTbpOzrxOfu")
	var builderCommitment, _ = tagged_base64.Parse("BUILDER_COMMITMENT~1yS-KEtL3oDZDBJdsW51Pd7zywIiHesBZsTbpOzrxOdZ")
	var blockMerkleTreeRoot, _ = tagged_base64.Parse("MERKLE_COMM~yB4_Aqa35_PoskgTpcCR1oVLh6BUdLHIs7erHKWi-usUAAAAAAAAAAEAAAAAAAAAJg")
	var feeMerkleTreeRoot, _ = tagged_base64.Parse("MERKLE_COMM~VJ9z239aP9GZDrHp3VxwPd_0l28Hc5KEAB1pFeCIxhYgAAAAAAAAAAIAAAAAAAAAdA")
	var blockInfo = common_types.L1BlockInfo{
		Number:    123,
		Timestamp: *common_types.NewU256().SetUint64(0x456),
		Hash:      common.Hash{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef, 0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef, 0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef, 0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef},
	}

	return v0_1.Header{
		ChainConfig: &v0_1.ResolvableChainConfig{
			ChainConfig: v0_1.EitherChainConfig{
				Left: &v0_1.ChainConfig{
					ChainId:      *common_types.NewU256().SetUint64(0x8a19).ToDecimal(),
					MaxBlockSize: *common_types.NewU256().SetUint64(10240).ToDecimal(),
					BaseFee:      *common_types.NewU256().SetUint64(0).ToDecimal(),
					FeeContract:  &common.Address{},
					FeeRecipient: common.Address{},
				},
			},
		},
		Height:    0,
		Timestamp: 789,
		L1Head:    124,
		NsTable: &common_types.NsTable{
			Bytes: []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		BlockMerkleTreeRoot: blockMerkleTreeRoot,
		FeeMerkleTreeRoot:   feeMerkleTreeRoot,
		BuilderCommitment:   builderCommitment,
		PayloadCommitment:   payloadCommitment,
		FeeInfo:             &common_types.FeeInfo{Account: common.HexToAddress("0xf39fd6e51aad88f6f4ce6ab8827279cfffb92266"), Amount: *common_types.NewU256().SetUint64(0).ToDecimal()},
		L1Finalized:         &blockInfo,
		BuilderSignature: &common_types.Signature{
			// R: common.HexToHash("0x1f92bab6350d4f33e04f9e9278d89f644d0abea16d6f882e91f87bec4e0ba53d"),
			// S: common.HexToHash("0x2067627270a89b06e7486c2c56fef0fee5f49a14b296a1cde580b0b40fa7430f"),
			R: "",
			S: "",
			V: uint64(27),
		},
	}
}
