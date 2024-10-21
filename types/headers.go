package types

import (
	"encoding/json"
	"fmt"
)

type Version struct {
	Major uint16 `json:"major"`
	Minor uint16 `json:"minor"`
}

func (v *Version) UnmarshalJSON(b []byte) error {
	// Use an alias type to avoid recusive calls of this function
	type Alias Version

	type Dec struct {
		Ver Alias `json:"Version"`
	}

	var dec Dec
	if err := json.Unmarshal(b, &dec); err != nil {
		return err
	}

	v.Major = dec.Ver.Major
	v.Minor = dec.Ver.Minor

	return nil
}

type HeaderInterface interface {
	Commit() Commitment
	Version() Version
	GetBlockHeight() uint64
	GetPayloadCommitment() *TaggedBase64
	GetBuilderCommitment() *TaggedBase64
	GetNsTable() *NsTable
	GetBlockMerkleTreeRoot() *TaggedBase64
	GetFeeMerkleTreeRoot() *TaggedBase64
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

type RawHeader struct {
	Version Version         `json:"version"`
	Fields  json.RawMessage `json:"fields"`
}

func (r *RawHeader) UnmarshalJSON(b []byte) error {
	type Dec struct {
		Version *Version         `json:"version"`
		Fields  *json.RawMessage `json:"fields"`
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
		var header Header
		// Failed to parse the version header. Legacy header
		if err = json.Unmarshal(data, &header); err != nil {
			return nil, err
		}

		return &header, nil
	}

	version := rawHeader.Version
	if version.Major == 0 && version.Minor == 2 {
		var header Header
		if err := json.Unmarshal(rawHeader.Fields, &header); err != nil {
			return nil, err
		}
		return &header, nil
	}

	if version.Major == 0 && version.Minor == 3 {
		var header Header0_3
		if err := json.Unmarshal(rawHeader.Fields, &header); err != nil {
			return nil, err
		}
		return &header, nil
	}

	return nil, fmt.Errorf("version error: %v", version)
}

type Header struct {
	ChainConfig         *ResolvableChainConfig `json:"chain_config"`
	Height              uint64                 `json:"height"`
	Timestamp           uint64                 `json:"timestamp"`
	L1Head              uint64                 `json:"l1_head"`
	L1Finalized         *L1BlockInfo           `json:"l1_finalized"           rlp:"nil"`
	PayloadCommitment   *TaggedBase64          `json:"payload_commitment"`
	BuilderCommitment   *TaggedBase64          `json:"builder_commitment"`
	NsTable             *NsTable               `json:"ns_table"`
	BlockMerkleTreeRoot *TaggedBase64          `json:"block_merkle_tree_root"`
	FeeMerkleTreeRoot   *TaggedBase64          `json:"fee_merkle_tree_root"`
	FeeInfo             *FeeInfo               `json:"fee_info"`
	BuilderSignature    *BuilderSignature      `json:"builder_signature"           rlp:"nil"`
}

func (h *Header) Version() Version {
	return Version{Major: 0, Minor: 1}
}

func (h *Header) GetBlockHeight() uint64 {
	return h.Height
}
func (h *Header) GetPayloadCommitment() *TaggedBase64 {
	return h.PayloadCommitment
}
func (h *Header) GetBuilderCommitment() *TaggedBase64 {
	return h.BuilderCommitment
}
func (h *Header) GetNsTable() *NsTable {
	return h.NsTable
}
func (h *Header) GetBlockMerkleTreeRoot() *TaggedBase64 {
	return h.BlockMerkleTreeRoot
}
func (h *Header) GetFeeMerkleTreeRoot() *TaggedBase64 {
	return h.FeeMerkleTreeRoot
}

func (h *Header) GetBuilderSignature() *BuilderSignature {
	return h.BuilderSignature
}

func (h *Header) UnmarshalJSON(b []byte) error {
	// Parse using pointers so we can distinguish between missing and default fields.
	type Dec struct {
		ChainConfig         **ResolvableChainConfig `json:"chain_config"`
		Height              *uint64                 `json:"height"`
		Timestamp           *uint64                 `json:"timestamp"`
		L1Head              *uint64                 `json:"l1_head"`
		L1Finalized         *L1BlockInfo            `json:"l1_finalized"           rlp:"nil"`
		PayloadCommitment   **TaggedBase64          `json:"payload_commitment"`
		BuilderCommitment   **TaggedBase64          `json:"builder_commitment"`
		NsTable             **NsTable               `json:"ns_table"`
		BlockMerkleTreeRoot **TaggedBase64          `json:"block_merkle_tree_root"`
		FeeMerkleTreeRoot   **TaggedBase64          `json:"fee_merkle_tree_root"`
		FeeInfo             **FeeInfo               `json:"fee_info"`
		BuilderSignature    *BuilderSignature       `json:"builder_signature"           rlp:"nil"`
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
	h.BuilderSignature = dec.BuilderSignature

	return nil
}

func (self *Header) Commit() Commitment {
	var l1FinalizedComm *Commitment
	if self.L1Finalized != nil {
		comm := self.L1Finalized.Commit()
		l1FinalizedComm = &comm
	}

	var builderSignatureCommitment *Commitment
	if self.BuilderSignature != nil {
		comm := self.BuilderSignature.Commit()
		builderSignatureCommitment = &comm
	}

	return NewRawCommitmentBuilder("BLOCK").
		Field("chain_config", self.ChainConfig.Commit()).
		Uint64Field("height", self.Height).
		Uint64Field("timestamp", self.Timestamp).
		Uint64Field("l1_head", self.L1Head).
		OptionalField("l1_finalized", l1FinalizedComm).
		ConstantString("payload_commitment").
		FixedSizeBytes(self.PayloadCommitment.Value()).
		ConstantString("builder_commitment").
		FixedSizeBytes(self.BuilderCommitment.Value()).
		Field("ns_table", self.NsTable.Commit()).
		VarSizeField("block_merkle_tree_root", self.BlockMerkleTreeRoot.Value()).
		VarSizeField("fee_merkle_tree_root", self.FeeMerkleTreeRoot.Value()).
		Field("fee_info", self.FeeInfo.Commit()).
		OptionalField("builder_signature", builderSignatureCommitment).
		Finalize()
}

type Header0_3 struct {
	ChainConfig         *ResolvableChainConfig `json:"chain_config"`
	Height              uint64                 `json:"height"`
	Timestamp           uint64                 `json:"timestamp"`
	L1Head              uint64                 `json:"l1_head"`
	L1Finalized         *L1BlockInfo           `json:"l1_finalized"           rlp:"nil"`
	PayloadCommitment   *TaggedBase64          `json:"payload_commitment"`
	BuilderCommitment   *TaggedBase64          `json:"builder_commitment"`
	NsTable             *NsTable               `json:"ns_table"`
	BlockMerkleTreeRoot *TaggedBase64          `json:"block_merkle_tree_root"`
	FeeMerkleTreeRoot   *TaggedBase64          `json:"fee_merkle_tree_root"`
	FeeInfo             *[]FeeInfo             `json:"fee_info"`
}

func (h *Header0_3) Version() Version {
	return Version{Major: 0, Minor: 3}
}

func (h *Header0_3) GetBlockHeight() uint64 {
	return h.Height
}
func (h *Header0_3) GetPayloadCommitment() *TaggedBase64 {
	return h.PayloadCommitment
}
func (h *Header0_3) GetBuilderCommitment() *TaggedBase64 {
	return h.BuilderCommitment
}
func (h *Header0_3) GetNsTable() *NsTable {
	return h.NsTable
}
func (h *Header0_3) GetBlockMerkleTreeRoot() *TaggedBase64 {
	return h.BlockMerkleTreeRoot
}
func (h *Header0_3) GetFeeMerkleTreeRoot() *TaggedBase64 {
	return h.FeeMerkleTreeRoot
}

func (h *Header0_3) UnmarshalJSON(b []byte) error {
	type Dec struct {
		ChainConfig         **ResolvableChainConfig `json:"chain_config"`
		Height              *uint64                 `json:"height"`
		Timestamp           *uint64                 `json:"timestamp"`
		L1Head              *uint64                 `json:"l1_head"`
		L1Finalized         *L1BlockInfo            `json:"l1_finalized"           rlp:"nil"`
		PayloadCommitment   **TaggedBase64          `json:"payload_commitment"`
		BuilderCommitment   **TaggedBase64          `json:"builder_commitment"`
		NsTable             **NsTable               `json:"ns_table"`
		BlockMerkleTreeRoot **TaggedBase64          `json:"block_merkle_tree_root"`
		FeeMerkleTreeRoot   **TaggedBase64          `json:"fee_merkle_tree_root"`
		FeeInfo             **[]FeeInfo             `json:"fee_info"`
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

func (self *Header0_3) Commit() Commitment {
	panic("unimplemented")
}
