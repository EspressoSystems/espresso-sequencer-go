package v0_3

import (
	"encoding/json"
	"fmt"

	common "github.com/EspressoSystems/espresso-sequencer-go/types/common"
)

type Header struct {
	ChainConfig         *common.ResolvableChainConfig `json:"chain_config"`
	Height              uint64                        `json:"height"`
	Timestamp           uint64                        `json:"timestamp"`
	L1Head              uint64                        `json:"l1_head"`
	L1Finalized         *common.L1BlockInfo           `json:"l1_finalized"           rlp:"nil"`
	PayloadCommitment   *common.TaggedBase64          `json:"payload_commitment"`
	BuilderCommitment   *common.TaggedBase64          `json:"builder_commitment"`
	NsTable             *common.NsTable               `json:"ns_table"`
	BlockMerkleTreeRoot *common.TaggedBase64          `json:"block_merkle_tree_root"`
	FeeMerkleTreeRoot   *common.TaggedBase64          `json:"fee_merkle_tree_root"`
	FeeInfo             *[]common.FeeInfo             `json:"fee_info"`
}

func (h *Header) Version() common.Version {
	return common.Version{Major: 0, Minor: 3}
}

func (h *Header) GetBlockHeight() uint64 {
	return h.Height
}
func (h *Header) GetPayloadCommitment() *common.TaggedBase64 {
	return h.PayloadCommitment
}
func (h *Header) GetBuilderCommitment() *common.TaggedBase64 {
	return h.BuilderCommitment
}
func (h *Header) GetNsTable() *common.NsTable {
	return h.NsTable
}
func (h *Header) GetBlockMerkleTreeRoot() *common.TaggedBase64 {
	return h.BlockMerkleTreeRoot
}
func (h *Header) GetFeeMerkleTreeRoot() *common.TaggedBase64 {
	return h.FeeMerkleTreeRoot
}

func (h *Header) UnmarshalJSON(b []byte) error {
	type Dec struct {
		ChainConfig         **common.ResolvableChainConfig `json:"chain_config"`
		Height              *uint64                        `json:"height"`
		Timestamp           *uint64                        `json:"timestamp"`
		L1Head              *uint64                        `json:"l1_head"`
		L1Finalized         *common.L1BlockInfo            `json:"l1_finalized"           rlp:"nil"`
		PayloadCommitment   **common.TaggedBase64          `json:"payload_commitment"`
		BuilderCommitment   **common.TaggedBase64          `json:"builder_commitment"`
		NsTable             **common.NsTable               `json:"ns_table"`
		BlockMerkleTreeRoot **common.TaggedBase64          `json:"block_merkle_tree_root"`
		FeeMerkleTreeRoot   **common.TaggedBase64          `json:"fee_merkle_tree_root"`
		FeeInfo             **[]common.FeeInfo             `json:"fee_info"`
	}

	var dec Dec
	if err := json.Unmarshal(b, &dec); err != nil {
		return err
	}

	if dec.Height == nil {
		return fmt.Errorf("Field height of type Header is required")
	}
	h.Height = *dec.Height

	if dec.Timestamp == nil {
		return fmt.Errorf("Field timestamp of type Header is required")
	}
	h.Timestamp = *dec.Timestamp

	if dec.L1Head == nil {
		return fmt.Errorf("Field l1_head of type Header is required")
	}
	h.L1Head = *dec.L1Head

	if dec.PayloadCommitment == nil {
		return fmt.Errorf("Field payload_commitment of type Header is required")
	}
	h.PayloadCommitment = *dec.PayloadCommitment

	if dec.BuilderCommitment == nil {
		return fmt.Errorf("Field builder_commitment of type Header is required")
	}
	h.BuilderCommitment = *dec.BuilderCommitment

	if dec.NsTable == nil {
		return fmt.Errorf("Field transactions_root of type Header is required")
	}
	h.NsTable = *dec.NsTable

	if dec.BlockMerkleTreeRoot == nil {
		return fmt.Errorf("Field block_merkle_tree_root of type Header is required")
	}
	h.BlockMerkleTreeRoot = *dec.BlockMerkleTreeRoot

	if dec.FeeMerkleTreeRoot == nil {
		return fmt.Errorf("Field fee_merkle_tree_root of type Header is required")
	}
	h.FeeMerkleTreeRoot = *dec.FeeMerkleTreeRoot

	if dec.FeeInfo == nil {
		return fmt.Errorf("Field fee_info of type Header is required")
	}
	h.FeeInfo = *dec.FeeInfo

	if dec.ChainConfig == nil {
		return fmt.Errorf("Field chain_info of type Header is required")
	}
	h.ChainConfig = *dec.ChainConfig

	h.L1Finalized = dec.L1Finalized
	return nil
}

func (self *Header) Commit() common.Commitment {
	panic("unimplemented")
}
