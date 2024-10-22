package types

import (
	"encoding/json"
	"io"
	"os"
	"strings"
	"testing"

	tagged_base64 "github.com/EspressoSystems/espresso-sequencer-go/tagged-base64"
	common_types "github.com/EspressoSystems/espresso-sequencer-go/types/common"
	"github.com/EspressoSystems/espresso-sequencer-go/types/v0/v0_1"
	common "github.com/ethereum/go-ethereum/common"

	"github.com/stretchr/testify/require"
)

var ReferenceL1BLockInfo common_types.L1BlockInfo = common_types.L1BlockInfo{
	Number:    123,
	Timestamp: *common_types.NewU256().SetUint64(0x456),
	Hash:      common.Hash{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef, 0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef, 0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef, 0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef},
}

var ReferenceNsTable common_types.NsTable = common_types.NsTable{
	Bytes: []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
}

var ReferenceChainConfig = &common_types.ResolvableChainConfig{
	ChainConfig: common_types.EitherChainConfig{
		Left: &common_types.ChainConfig{
			ChainId:      *common_types.NewU256().SetUint64(0x8a19).ToDecimal(),
			MaxBlockSize: *common_types.NewU256().SetUint64(10240).ToDecimal(),
			BaseFee:      *common_types.NewU256().SetUint64(0).ToDecimal(),
			FeeContract:  &common.Address{},
			FeeRecipient: common.Address{},
		},
	},
}

var ReferencePayloadCommitment, _ = tagged_base64.Parse("HASH~1yS-KEtL3oDZDBJdsW51Pd7zywIiHesBZsTbpOzrxOfu")
var ReferenceBuilderCommitment, _ = tagged_base64.Parse("BUILDER_COMMITMENT~1yS-KEtL3oDZDBJdsW51Pd7zywIiHesBZsTbpOzrxOdZ")
var ReferenceBlockMerkleTreeRoot, _ = tagged_base64.Parse("MERKLE_COMM~yB4_Aqa35_PoskgTpcCR1oVLh6BUdLHIs7erHKWi-usUAAAAAAAAAAEAAAAAAAAAJg")
var ReferenceFeeMerkleTreeRoot, _ = tagged_base64.Parse("MERKLE_COMM~VJ9z239aP9GZDrHp3VxwPd_0l28Hc5KEAB1pFeCIxhYgAAAAAAAAAAIAAAAAAAAAdA")

var ReferenceHeader v0_1.Header = v0_1.Header{
	ChainConfig:         ReferenceChainConfig,
	Height:              42,
	Timestamp:           789,
	L1Head:              124,
	L1Finalized:         &ReferenceL1BLockInfo,
	PayloadCommitment:   ReferencePayloadCommitment,
	BuilderCommitment:   ReferenceBuilderCommitment,
	NsTable:             &ReferenceNsTable,
	BlockMerkleTreeRoot: ReferenceBlockMerkleTreeRoot,
	FeeMerkleTreeRoot:   ReferenceFeeMerkleTreeRoot,
	FeeInfo:             &common_types.FeeInfo{Account: common.HexToAddress("0xf39fd6e51aad88f6f4ce6ab8827279cfffb92266"), Amount: *common_types.NewU256().SetUint64(0).ToDecimal()},
	BuilderSignature: &common_types.BuilderSignature{
		R: common.HexToHash("0x1f92bab6350d4f33e04f9e9278d89f644d0abea16d6f882e91f87bec4e0ba53d"),
		S: common.HexToHash("0x2067627270a89b06e7486c2c56fef0fee5f49a14b296a1cde580b0b40fa7430f"),
		V: 27,
	},
}

func TestVersion(t *testing.T) {
	s := `{"Version":{"major":0,"minor":3}}`
	data := []byte(s)
	var v common_types.Version
	err := json.Unmarshal(data, &v)
	if err != nil {
		t.Fatalf("Failed to marshal JSON: %v", err)
	}

	if !(v.Major == 0 && v.Minor == 3) {
		t.Fatalf("%v", v)
	}
}

func TestHeader0_1(t *testing.T) {
	file, err := os.Open("./test-data/header0_1.json")
	if err != nil {
		t.Fatal("failed to open file:", err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		t.Fatal("Error reading file:", err)
	}

	var headerImpl HeaderImpl
	err = json.Unmarshal(data, &headerImpl)
	if err != nil {
		t.Fatal("Error unmarshaling:", err)
	}

	header := headerImpl.Header

	if header.Version().Major != 0 || header.Version().Minor != 1 {
		t.Fatal("Wrong version", header.Version())
	}

	if header.GetBlockHeight() != 42 {
		t.Fatal("Wrong block height", header.GetBlockHeight())
	}

}

func TestHeader0_3(t *testing.T) {
	file, err := os.Open("./test-data/header0_3.json")
	if err != nil {
		t.Fatal("failed to open file:", err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		t.Fatal("Error reading file:", err)
	}

	var headerImpl HeaderImpl
	err = json.Unmarshal(data, &headerImpl)
	if err != nil {
		t.Fatal("Error unmarshaling:", err)
	}

	header := headerImpl.Header

	if header.Version().Major != 0 || header.Version().Minor != 3 {
		t.Fatal("Wrong version", header.Version())
	}

	if header.GetBlockHeight() != 1 {
		t.Fatal("Wrong block height", header.GetBlockHeight())
	}
}

func removeWhitespace(s string) string {
	// Split the string on whitespace then concatenate the segments
	return strings.Join(strings.Fields(s), "")
}

func CheckJsonRequiredFields[T any](t *testing.T, data []byte, fields ...string) {
	// Parse the JSON object into a map so we can selectively delete fields.
	var obj map[string]json.RawMessage
	if err := json.Unmarshal(data, &obj); err != nil {
		t.Fatalf("failed to unmarshal JSON: %v", err)
	}

	for _, field := range fields {
		data, err := json.Marshal(withoutKey(obj, field))
		require.Nil(t, err, "failed to marshal JSON")

		var dec T
		err = json.Unmarshal(data, &dec)
		require.NotNil(t, err, "unmarshalling without required field %s should fail", field)
	}
}

func withoutKey[K comparable, V any](m map[K]V, key K) map[K]V {
	copied := make(map[K]V)
	for k, v := range m {
		if k != key {
			copied[k] = v
		}
	}
	return copied
}

func TestEspressoTypesHeaderJson(t *testing.T) {
	data := []byte(removeWhitespace(`{
		"chain_config": {
			"chain_config": {
			  "Left": {
				"chain_id": "35353",
				"max_block_size": "10240",
				"base_fee": "0",
				"fee_contract": "0x0000000000000000000000000000000000000000",
				"fee_recipient": "0x0000000000000000000000000000000000000000"
			  }
			}
		},
		"height": 42,
		"timestamp": 789,
		"l1_head": 124,
		"l1_finalized": {
			"number": 123,
			"timestamp": "0x456",
			"hash": "0x0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef"
		},
		"payload_commitment": "HASH~1yS-KEtL3oDZDBJdsW51Pd7zywIiHesBZsTbpOzrxOfu",
		"builder_commitment": "BUILDER_COMMITMENT~1yS-KEtL3oDZDBJdsW51Pd7zywIiHesBZsTbpOzrxOdZ",
		"ns_table": {
			"bytes":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"
		},
		"block_merkle_tree_root": "MERKLE_COMM~yB4_Aqa35_PoskgTpcCR1oVLh6BUdLHIs7erHKWi-usUAAAAAAAAAAEAAAAAAAAAJg",
		"fee_merkle_tree_root": "MERKLE_COMM~VJ9z239aP9GZDrHp3VxwPd_0l28Hc5KEAB1pFeCIxhYgAAAAAAAAAAIAAAAAAAAAdA",
		"fee_info":{"account":"0xf39fd6e51aad88f6f4ce6ab8827279cfffb92266","amount":"0"},
		"builder_signature":{"r":"0x1f92bab6350d4f33e04f9e9278d89f644d0abea16d6f882e91f87bec4e0ba53d","s":"0x2067627270a89b06e7486c2c56fef0fee5f49a14b296a1cde580b0b40fa7430f","v":27}
	}`))

	// Check encoding.
	encoded, err := json.Marshal(ReferenceHeader)
	if err != nil {
		t.Fatalf("Failed to marshal JSON: %v", err)
	}
	require.Equal(t, encoded, data)

	// Check decoding
	var decoded v0_1.Header
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", err)
	}
	require.Equal(t, decoded, ReferenceHeader)

	CheckJsonRequiredFields[v0_1.Header](t, data, "height", "timestamp", "l1_head", "payload_commitment", "builder_commitment", "block_merkle_tree_root", "fee_merkle_tree_root", "fee_info", "chain_config")
}

func TestEspressoTypesHeaderCommit(t *testing.T) {
	require.Equal(t, ReferenceHeader.Commit(), common_types.Commitment{231, 29, 239, 71, 22, 101, 198, 223, 100, 115, 197, 203, 125, 6, 172, 149, 141, 226, 79, 254, 84, 28, 152, 122, 37, 222, 143, 27, 192, 81, 80, 57})
}
