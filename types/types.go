package types

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/big"

	tagged_base64 "github.com/EspressoSystems/espresso-sequencer-go/tagged-base64"

	"github.com/ethereum/go-ethereum/common"
)

type TaggedBase64 = tagged_base64.TaggedBase64

type VidCommon = json.RawMessage

type VidCommonQueryData struct {
	Height      uint64        `json:"height"`
	BlockHash   *TaggedBase64 `json:"block_hash"`
	PayloadHash *TaggedBase64 `json:"payload_hash"`
	Common      VidCommon     `json:"common"`
}

type TransactionQueryData struct {
	Transaction Transaction     `json:"transaction"`
	Hash        *TaggedBase64   `json:"hash"`
	Index       uint64          `json:"index"`
	Proof       json.RawMessage `json:"proof"`
	BlockHash   *TaggedBase64   `json:"block_hash"`
	BlockHeight uint64          `json:"block_height"`
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

func (self *Header) Commit() Commitment {
	var l1FinalizedComm *Commitment
	if self.L1Finalized != nil {
		comm := self.L1Finalized.Commit()
		l1FinalizedComm = &comm
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
		Finalize()
}

type ResolvableChainConfig struct {
	ChainConfig EitherChainConfig `json:"chain_config"`
}

func (self *ResolvableChainConfig) Commit() Commitment {
	config := self.ChainConfig
	if config.Left != nil {
		return config.Left.Commit()
	}
	if config.Right != nil {
		right := *config.Right
		r := (*right).Value()
		bytes := [32]byte{}
		copy(bytes[:], r)
		return Commitment(bytes)
	}

	// It shouldn't happen
	return Commitment{}
}

type EitherChainConfig struct {
	Left  *ChainConfig   `json:"Left"`
	Right **TaggedBase64 `json:"Right"`
}

func (i *EitherChainConfig) UnmarshalJSON(b []byte) error {
	type Dec struct {
		Left  *ChainConfig   `json:"Left"`
		Right **TaggedBase64 `json:"Right"`
	}
	var dec Dec
	if err := json.Unmarshal(b, &dec); err != nil {
		return err
	}
	if dec.Left != nil {
		i.Left = dec.Left
	}

	if dec.Right != nil {
		i.Right = dec.Right
	}

	if i.Left == nil && i.Right == nil {
		return fmt.Errorf("either Left or Right variant for EitherChainConfig is required")
	}

	return nil
}

func (i *EitherChainConfig) MarshalJSON() ([]byte, error) {
	type Left struct {
		Left *ChainConfig `json:"Left"`
	}
	type Right struct {
		Right **TaggedBase64 `json:"Right"`
	}

	if i.Left != nil {
		return json.Marshal(Left{Left: i.Left})
	}

	return json.Marshal(Right{Right: i.Right})
}

type ChainConfig struct {
	ChainId      U256Decimal     `json:"chain_id"`
	MaxBlockSize U256Decimal     `json:"max_block_size"`
	BaseFee      U256Decimal     `json:"base_fee"`
	FeeContract  *common.Address `json:"fee_contract"`
	FeeRecipient common.Address  `json:"fee_recipient"`
}

func (self *ChainConfig) Commit() Commitment {
	builder := NewRawCommitmentBuilder("CHAIN_CONFIG").
		Uint256Field("chain_id", self.ChainId.ToU256()).
		Uint64Field("max_block_size", self.MaxBlockSize.Uint64()).
		Uint256Field("base_fee", self.BaseFee.ToU256()).
		FixedSizeField("fee_recipient", self.FeeRecipient.Bytes())
	if self.FeeContract == nil {
		return builder.Finalize()
	}
	return builder.Uint64Field("fee_contract", 1).FixedSizeBytes(self.FeeContract.Bytes()).Finalize()
}

type L1BlockInfo struct {
	Number    uint64      `json:"number"`
	Timestamp U256        `json:"timestamp"`
	Hash      common.Hash `json:"hash"`
}

func (i *L1BlockInfo) UnmarshalJSON(b []byte) error {
	// Parse using pointers so we can distinguish between missing and default fields.
	type Dec struct {
		Number    *uint64      `json:"number"`
		Timestamp *U256        `json:"timestamp"`
		Hash      *common.Hash `json:"hash"`
	}

	var dec Dec
	if err := json.Unmarshal(b, &dec); err != nil {
		return err
	}

	if dec.Number == nil {
		return fmt.Errorf("Field number of type L1BlockInfo is required")
	}
	i.Number = *dec.Number

	if dec.Timestamp == nil {
		return fmt.Errorf("Field timestamp of type L1BlockInfo is required")
	}
	i.Timestamp = *dec.Timestamp

	if dec.Hash == nil {
		return fmt.Errorf("Field hash of type L1BlockInfo is required")
	}
	i.Hash = *dec.Hash

	return nil
}

func (self *L1BlockInfo) Commit() Commitment {
	return NewRawCommitmentBuilder("L1BLOCK").
		Uint64Field("number", self.Number).
		Uint256Field("timestamp", &self.Timestamp).
		FixedSizeField("hash", self.Hash[:]).
		Finalize()
}

type NsTable struct {
	Bytes Bytes `json:"bytes"`
}

type NamespaceProof = json.RawMessage

type BlockMerkleRoot = Commitment

type HotShotBlockMerkleProof struct {
	Proof json.RawMessage `json:"proof"`
}

// Validates a block merkle proof, returning the validated HotShot block height. This is mocked until we have real
// merkle tree snapshot support.
func (p *HotShotBlockMerkleProof) Verify(root BlockMerkleRoot) (uint64, error) {
	return 0, nil
}

func (r *NsTable) UnmarshalJSON(b []byte) error {
	// Parse using pointers so we can distinguish between missing and default fields.
	type Dec struct {
		Bytes *Bytes `json:"bytes"`
	}

	var dec Dec
	if err := json.Unmarshal(b, &dec); err != nil {
		return err
	}

	if dec.Bytes == nil {
		return fmt.Errorf("Field root of type RawPayload is required")
	}
	r.Bytes = *dec.Bytes

	return nil
}

func (self *NsTable) Commit() Commitment {
	return NewRawCommitmentBuilder("NSTABLE").
		VarSizeBytes(self.Bytes).
		Finalize()
}

type Transaction struct {
	Namespace uint64 `json:"namespace"`
	Payload   Bytes  `json:"payload"`
}

func (self *Transaction) Commit() Commitment {
	return NewRawCommitmentBuilder("Transaction").
		Uint64Field("namespace", self.Namespace).
		VarSizeBytes(self.Payload).
		Finalize()
}

func (t *Transaction) UnmarshalJSON(b []byte) error {
	// Parse using pointers so we can distinguish between missing and default fields.
	type Dec struct {
		Namespace *uint64 `json:"namespace"`
		Payload   *Bytes  `json:"payload"`
	}

	var dec Dec
	if err := json.Unmarshal(b, &dec); err != nil {
		return err
	}

	if dec.Namespace == nil {
		return fmt.Errorf("Field vm of type Transaction is required")
	}
	t.Namespace = *dec.Namespace

	if dec.Payload == nil {
		return fmt.Errorf("Field payload of type Transaction is required")
	}
	t.Payload = *dec.Payload

	return nil
}

// A bytes type which serializes to JSON as an array, rather than a base64 string. This ensures
// compatibility with the Espresso APIs.
type Bytes []byte

func (b Bytes) MarshalJSON() ([]byte, error) {
	s := base64.StdEncoding.EncodeToString(b)
	return json.Marshal(s)
}

func (b *Bytes) UnmarshalJSON(in []byte) error {
	var s string
	if err := json.Unmarshal(in, &s); err != nil {
		return err
	}
	bytes, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return err
	}
	*b = bytes
	return nil
}

// A readable decimal format for U256. Please use the struct `U256` to initialize
// the number first and use the `ToDecimal` to convert.
type U256Decimal struct {
	big.Int
}

func (i U256Decimal) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.Text(10))
}

func (i *U256Decimal) UnmarshalJSON(in []byte) error {
	var s string
	if err := json.Unmarshal(in, &s); err != nil {
		return err
	}
	if _, err := fmt.Sscanf(s, "%d", &i.Int); err != nil {
		return err
	}
	return nil
}

func (i *U256Decimal) ToU256() *U256 {
	return &U256{i.Int}
}

// A BigInt type which serializes to JSON a a hex string. This ensures compatibility with the
// Espresso APIs.
type U256 struct {
	big.Int
}

func NewU256() *U256 {
	return new(U256)
}

func (i U256) Equal(other U256) bool {
	return i.Int.Cmp(&other.Int) == 0
}

func (i *U256) SetBigInt(n *big.Int) *U256 {
	i.Int.Set(n)
	return i
}

func (i *U256) SetUint64(n uint64) *U256 {
	i.Int.SetUint64(n)
	return i
}

func (i *U256) SetBytes(buf [32]byte) *U256 {
	i.Int.SetBytes(buf[:])
	return i
}

func (i U256) MarshalJSON() ([]byte, error) {
	return json.Marshal(fmt.Sprintf("0x%s", i.Text(16)))
}

func (i *U256) UnmarshalJSON(in []byte) error {
	var s string
	if err := json.Unmarshal(in, &s); err != nil {
		return err
	}
	if _, err := fmt.Sscanf(s, "0x%x", &i.Int); err != nil {
		return err
	}
	return nil
}

func (i *U256) ToDecimal() *U256Decimal {
	return &U256Decimal{i.Int}
}

type FeeInfo struct {
	Account common.Address `json:"account"`
	Amount  U256Decimal    `json:"amount"`
}

func (self *FeeInfo) Commit() Commitment {
	return NewRawCommitmentBuilder("FEE_INFO").
		FixedSizeField("account", self.Account.Bytes()).
		Uint256Field("amount", self.Amount.ToU256()).
		Finalize()
}

// For some intricate reasons, rollups need to build a dummy header as a placeholder. However, an empty Header can be serialized
// but can not be deserialized because of the checks. Thus we provide this dummy header as a workaround.
func GetDummyHeader() Header {
	var payloadCommitment, _ = tagged_base64.Parse("HASH~1yS-KEtL3oDZDBJdsW51Pd7zywIiHesBZsTbpOzrxOfu")
	var builderCommitment, _ = tagged_base64.Parse("BUILDER_COMMITMENT~1yS-KEtL3oDZDBJdsW51Pd7zywIiHesBZsTbpOzrxOdZ")
	var blockMerkleTreeRoot, _ = tagged_base64.Parse("MERKLE_COMM~yB4_Aqa35_PoskgTpcCR1oVLh6BUdLHIs7erHKWi-usUAAAAAAAAAAEAAAAAAAAAJg")
	var feeMerkleTreeRoot, _ = tagged_base64.Parse("MERKLE_COMM~VJ9z239aP9GZDrHp3VxwPd_0l28Hc5KEAB1pFeCIxhYgAAAAAAAAAAIAAAAAAAAAdA")
	var blockInfo = L1BlockInfo{
		Number:    123,
		Timestamp: *NewU256().SetUint64(0x456),
		Hash:      common.Hash{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef, 0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef, 0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef, 0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef},
	}

	return Header{
		ChainConfig: &ResolvableChainConfig{
			EitherChainConfig{
				Left: &ChainConfig{
					ChainId:      *NewU256().SetUint64(0x8a19).ToDecimal(),
					MaxBlockSize: *NewU256().SetUint64(10240).ToDecimal(),
					BaseFee:      *NewU256().SetUint64(0).ToDecimal(),
					FeeContract:  &common.Address{},
					FeeRecipient: common.Address{},
				},
			},
		},
		Height:    0,
		Timestamp: 789,
		L1Head:    124,
		NsTable: &NsTable{
			Bytes: []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		BlockMerkleTreeRoot: blockMerkleTreeRoot,
		FeeMerkleTreeRoot:   feeMerkleTreeRoot,
		BuilderCommitment:   builderCommitment,
		PayloadCommitment:   payloadCommitment,
		FeeInfo:             &FeeInfo{Account: common.HexToAddress("0xf39fd6e51aad88f6f4ce6ab8827279cfffb92266"), Amount: *NewU256().SetUint64(0).ToDecimal()},
		L1Finalized:         &blockInfo,
	}
}
