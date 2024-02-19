package types

import (
	"encoding/json"
	"strings"
	"testing"

	tagged_base64 "github.com/EspressoSystems/espresso-sequencer-go/tagged-base64"

	"github.com/ethereum/go-ethereum/common"

	"github.com/stretchr/testify/require"
)

func removeWhitespace(s string) string {
	// Split the string on whitespace then concatenate the segments
	return strings.Join(strings.Fields(s), "")
}

// Reference data taken from the reference sequencer implementation
// (https://github.com/EspressoSystems/espresso-sequencer/blob/main/data)

var ReferenceL1BLockInfo L1BlockInfo = L1BlockInfo{
	Number:    123,
	Timestamp: *NewU256().SetUint64(0x456),
	Hash:      common.Hash{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef, 0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef, 0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef, 0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef},
}

var ReferenceNsTable NsTable = NsTable{
	RawPayload: []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
}

var ReferencePayloadCommitment, _ = tagged_base64.Parse("HASH~1yS-KEtL3oDZDBJdsW51Pd7zywIiHesBZsTbpOzrxOfu")
var ReferenceBlockMerkleTreeRoot, _ = tagged_base64.Parse("MERKLE_COMM~yB4_Aqa35_PoskgTpcCR1oVLh6BUdLHIs7erHKWi-usUAAAAAAAAAAEAAAAAAAAAJg")
var ReferenceFeeMerkleTreeRoot, _ = tagged_base64.Parse("MERKLE_COMM~VJ9z239aP9GZDrHp3VxwPd_0l28Hc5KEAB1pFeCIxhYgAAAAAAAAAAIAAAAAAAAAdA")

var ReferenceHeader Header = Header{
	Height:              42,
	Timestamp:           789,
	L1Head:              124,
	L1Finalized:         &ReferenceL1BLockInfo,
	NsTable:             &ReferenceNsTable,
	PayloadCommitment:   ReferencePayloadCommitment,
	BlockMerkleTreeRoot: ReferenceBlockMerkleTreeRoot,
	FeeMerkleTreeRoot:   ReferenceFeeMerkleTreeRoot,
}

func TestEspressoTypesL1BLockInfoJson(t *testing.T) {
	data := []byte(removeWhitespace(`{
		"number": 123,
		"timestamp": "0x456",
		"hash": "0x0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef"
	}`))

	// Check encoding.
	encoded, err := json.Marshal(ReferenceL1BLockInfo)
	if err != nil {
		t.Fatalf("Failed to marshal JSON: %v", err)
	}
	require.Equal(t, encoded, data)

	// Check decoding
	var decoded L1BlockInfo
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", err)
	}
	require.Equal(t, decoded, ReferenceL1BLockInfo)

	CheckJsonRequiredFields[L1BlockInfo](t, data, "number", "timestamp", "hash")
}

func TestEspressoTypesHeaderJson(t *testing.T) {
	data := []byte(removeWhitespace(`{
		"height": 42,
		"timestamp": 789,
		"l1_head": 124,
		"l1_finalized": {
			"number": 123,
			"timestamp": "0x456",
			"hash": "0x0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef"
		},
		"ns_table": {
			"raw_payload":[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0]
		},
		"payload_commitment": "HASH~1yS-KEtL3oDZDBJdsW51Pd7zywIiHesBZsTbpOzrxOfu",
		"block_merkle_tree_root": "MERKLE_COMM~yB4_Aqa35_PoskgTpcCR1oVLh6BUdLHIs7erHKWi-usUAAAAAAAAAAEAAAAAAAAAJg",
		"fee_merkle_tree_root": "MERKLE_COMM~VJ9z239aP9GZDrHp3VxwPd_0l28Hc5KEAB1pFeCIxhYgAAAAAAAAAAIAAAAAAAAAdA"
	}`))

	// Check encoding.
	encoded, err := json.Marshal(ReferenceHeader)
	if err != nil {
		t.Fatalf("Failed to marshal JSON: %v", err)
	}
	require.Equal(t, encoded, data)

	// Check decoding
	var decoded Header
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", err)
	}
	require.Equal(t, decoded, ReferenceHeader)

	CheckJsonRequiredFields[Header](t, data, "height", "timestamp", "l1_head", "payload_commitment", "block_merkle_tree_root", "fee_merkle_tree_root")
}

func TestEspressoTransactionJson(t *testing.T) {
	data := []byte(removeWhitespace(`{
		"vm": 0,
		"payload": [1,2,3,4,5]
	}`))
	tx := Transaction{
		Vm:      0,
		Payload: []byte{1, 2, 3, 4, 5},
	}

	// Check encoding.
	encoded, err := json.Marshal(tx)
	if err != nil {
		t.Fatalf("Failed to marshal JSON: %v", err)
	}
	require.Equal(t, encoded, data)

	// Check decoding
	var decoded Transaction
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", err)
	}
	require.Equal(t, decoded, tx)

	CheckJsonRequiredFields[Transaction](t, data, "vm", "payload")
}

// Commitment tests ported from the reference sequencer implementation
// (https://github.com/EspressoSystems/espresso-sequencer/blob/main/sequencer/src/block.rs)

func TestEspressoTypesL1BlockInfoCommit(t *testing.T) {
	require.Equal(t, ReferenceL1BLockInfo.Commit(), Commitment{224, 122, 115, 150, 226, 202, 216, 139, 51, 221, 23, 79, 54, 243, 107, 12, 12, 144, 113, 99, 133, 217, 207, 73, 120, 182, 115, 84, 210, 230, 126, 148})
}

func TestEspressoTypesNsTable(t *testing.T) {
	require.Equal(t, ReferenceNsTable.Commit(), Commitment{24, 191, 165, 16, 16, 48, 53, 144, 229, 119, 16, 233, 201, 36, 89, 64, 40, 77, 158, 105, 253, 188, 220, 221, 32, 2, 252, 91, 209, 13, 58, 232})
}

func TestEspressoTypesHeaderCommit(t *testing.T) {
	require.Equal(t, ReferenceHeader.Commit(), Commitment{19, 245, 91, 19, 221, 210, 135, 172, 20, 94, 23, 207, 136, 96, 55, 55, 54, 123, 194, 75, 253, 27, 178, 32, 197, 202, 228, 155, 161, 34, 164, 35})
}

func TestEspressoCommitmentFromU256TrailingZero(t *testing.T) {
	comm := Commitment{209, 146, 197, 195, 145, 148, 17, 211, 52, 72, 28, 120, 88, 182, 204, 206, 77, 36, 56, 35, 3, 143, 77, 186, 69, 233, 104, 30, 90, 105, 48, 0}
	roundTrip, err := CommitmentFromUint256(comm.Uint256())
	require.Nil(t, err)
	require.Equal(t, comm, roundTrip)
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
