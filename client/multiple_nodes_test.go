package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"testing"
	"time"

	types "github.com/EspressoSystems/espresso-sequencer-go/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockClient is a mock implementation of the Client interface
type MockClient struct {
	mock.Mock
}

func (m *MockClient) FetchRawHeaderByHeight(ctx context.Context, height uint64) (json.RawMessage, error) {
	args := m.Called(ctx, height)
	return args.Get(0).(json.RawMessage), args.Error(1)
}

func TestFetchWithMajority(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	mockNode1 := new(MockClient)
	mockNode2 := new(MockClient)
	mockNode3 := new(MockClient)

	// Simulate a scenario where two nodes return the same result and one returns a different result
	mockNode1.On("FetchRawHeaderByHeight", ctx, uint64(1)).Return(json.RawMessage(`{"data":"value1"}`), nil)
	mockNode2.On("FetchRawHeaderByHeight", ctx, uint64(1)).Return(json.RawMessage(`{"data":"value1"}`), nil)
	mockNode3.On("FetchRawHeaderByHeight", ctx, uint64(1)).Return(json.RawMessage(`{"data":"value2"}`), nil)

	nodes := []*MockClient{mockNode1, mockNode2, mockNode3}

	result, err := FetchWithMajority(ctx, nodes, func(node *MockClient) (json.RawMessage, error) {
		return node.FetchRawHeaderByHeight(ctx, 1)
	})

	assert.NoError(t, err)
	assert.Equal(t, json.RawMessage(`{"data":"value1"}`), result)

	// Simulate a scenario where no majority is reached
	mockNode1.On("FetchRawHeaderByHeight", ctx, uint64(2)).Return(json.RawMessage(`{"data":"value1"}`), nil)
	mockNode2.On("FetchRawHeaderByHeight", ctx, uint64(2)).Return(json.RawMessage(`{"data":"value2"}`), nil)
	mockNode3.On("FetchRawHeaderByHeight", ctx, uint64(2)).Return(json.RawMessage(`{"data":"value3"}`), nil)

	_, err = FetchWithMajority(ctx, nodes, func(node *MockClient) (json.RawMessage, error) {
		return node.FetchRawHeaderByHeight(ctx, 2)
	})

	assert.Error(t, err)
	assert.Equal(t, "no majority consensus reached", err.Error())

	// Simulate a scenario where all nodes return an error
	mockNode1.On("FetchRawHeaderByHeight", ctx, uint64(3)).Return(json.RawMessage{}, errors.New("error"))
	mockNode2.On("FetchRawHeaderByHeight", ctx, uint64(3)).Return(json.RawMessage{}, errors.New("error"))
	mockNode3.On("FetchRawHeaderByHeight", ctx, uint64(3)).Return(json.RawMessage{}, errors.New("error"))

	_, err = FetchWithMajority(ctx, nodes, func(node *MockClient) (json.RawMessage, error) {
		return node.FetchRawHeaderByHeight(ctx, 3)
	})

	assert.Error(t, err)

	// Simulate a scenario where the majority returns same result but not the same order
	mockNode1.On("FetchRawHeaderByHeight", ctx, uint64(4)).Return(json.RawMessage(`{"key": "key", "data":"value1"}`), nil)
	mockNode2.On("FetchRawHeaderByHeight", ctx, uint64(4)).Return(json.RawMessage(`{"data":"value1", "key": "key"}`), nil)
	mockNode3.On("FetchRawHeaderByHeight", ctx, uint64(4)).Return(json.RawMessage(`{"key": "key", "data":"value2"}`), nil)

	result, err = FetchWithMajority(ctx, nodes, func(node *MockClient) (json.RawMessage, error) {
		return node.FetchRawHeaderByHeight(ctx, 4)
	})

	assert.NoError(t, err)
	expected, err := hashNormalizedJSON(json.RawMessage(`{"data":"value1", "key": "key"}`))
	assert.NoError(t, err)
	actual, err := hashNormalizedJSON(result)
	assert.NoError(t, err)
	// keys have been sorted
	assert.Equal(t, expected, actual)

	// Simulate a scenario where only a single node is available
	newNodes := []*MockClient{mockNode1}
	result, err = FetchWithMajority(ctx, newNodes, func(node *MockClient) (json.RawMessage, error) {
		return node.FetchRawHeaderByHeight(ctx, 1)
	})
	assert.NoError(t, err)
	assert.Equal(t, json.RawMessage(`{"data":"value1"}`), result)

	// Simulate a scenario where the response is an array
	mockNode1.On("FetchRawHeaderByHeight", ctx, uint64(5)).Return(json.RawMessage(`[{"data":"value1"}, {"data":"value2"}]`), nil)
	mockNode2.On("FetchRawHeaderByHeight", ctx, uint64(5)).Return(json.RawMessage(`[{"data":"value1"}, {"data":"value2"}]`), nil)
	mockNode3.On("FetchRawHeaderByHeight", ctx, uint64(5)).Return(json.RawMessage(`[{"data":"value1"}, {"data":"value2"}, {"data":"value3"}]`), nil)

	result, err = FetchWithMajority(ctx, nodes, func(node *MockClient) (json.RawMessage, error) {
		return node.FetchRawHeaderByHeight(ctx, 5)
	})
	assert.NoError(t, err)
	assert.Equal(t, json.RawMessage(`[{"data":"value1"}, {"data":"value2"}]`), result)

	// Simulate a scenario where some of nodes are not responding
	mockNode1.On("FetchRawHeaderByHeight", ctx, uint64(6)).Return(json.RawMessage(`[{"data":"value1"}, {"data":"value2"}]`), nil)
	mockNode2.On("FetchRawHeaderByHeight", ctx, uint64(6)).Return(json.RawMessage{}, errors.New("error"))
	mockNode3.On("FetchRawHeaderByHeight", ctx, uint64(6)).Return(json.RawMessage(`[{"data":"value1"}, {"data":"value2"}]`), nil)

	result, err = FetchWithMajority(ctx, nodes, func(node *MockClient) (json.RawMessage, error) {
		return node.FetchRawHeaderByHeight(ctx, 6)
	})
	assert.NoError(t, err)
	assert.Equal(t, json.RawMessage(`[{"data":"value1"}, {"data":"value2"}]`), result)

	// Simulate a scenario where response is nested type
	mockNode1.On("FetchRawHeaderByHeight", ctx, uint64(7)).Return(json.RawMessage(`{"data":{"key":"value", "key2":"[1,2,3,4]"}}`), nil)
	mockNode2.On("FetchRawHeaderByHeight", ctx, uint64(7)).Return(json.RawMessage(`{"data":{"key":"value", "key2":"[1,2,3]"}}`), nil)
	mockNode3.On("FetchRawHeaderByHeight", ctx, uint64(7)).Return(json.RawMessage(`{"data":{"key":"value", "key2":"[1,2,3]"}}`), nil)

	result, err = FetchWithMajority(ctx, nodes, func(node *MockClient) (json.RawMessage, error) {
		return node.FetchRawHeaderByHeight(ctx, 7)
	})
	assert.NoError(t, err)
	assert.Equal(t, json.RawMessage(`{"data":{"key":"value", "key2":"[1,2,3]"}}`), result)

	// Test with the mock header data
	header1 := types.HeaderImpl{Header: getHeaderFromTestFile("../types/test-data/header0_2.json", t)}
	header2 := types.HeaderImpl{Header: getHeaderFromTestFile("../types/test-data/header0_3.json", t)}
	data1, err := json.Marshal(header1)
	assert.NoError(t, err)
	data2, err := json.MarshalIndent(header1, "", "  ") // With indent
	assert.NoError(t, err)
	data3, err := json.Marshal(header2)
	assert.NoError(t, err)
	mockNode1.On("FetchRawHeaderByHeight", ctx, uint64(8)).Return(json.RawMessage(data1), nil)
	mockNode2.On("FetchRawHeaderByHeight", ctx, uint64(8)).Return(json.RawMessage(data2), nil)
	mockNode3.On("FetchRawHeaderByHeight", ctx, uint64(8)).Return(json.RawMessage(data3), nil)
	result, err = FetchWithMajority(ctx, nodes, func(node *MockClient) (json.RawMessage, error) {
		return node.FetchRawHeaderByHeight(ctx, 8)
	})
	assert.NoError(t, err)
	var resultHeader types.HeaderImpl
	err = json.Unmarshal(result, &resultHeader)
	fmt.Println(resultHeader.Header.Version())
	assert.NoError(t, err)
	assert.Equal(t, header1.Header.Commit(), resultHeader.Header.Commit())
}

func TestApiWithSingleEspressoDevNode(t *testing.T) {
	ctx := context.Background()
	cleanup := runEspresso()
	defer cleanup()

	err := waitForEspressoNode(ctx)
	if err != nil {
		t.Fatal("failed to start espresso dev node", err)
	}

	client := NewMultipleNodesClient([]string{"http://localhost:21000"})

	blockHeight, err := client.FetchLatestBlockHeight(ctx)
	if err != nil {
		t.Fatal("failed to fetch block height")
	}

	_, err = client.FetchHeaderByHeight(ctx, blockHeight-1)
	if err != nil {
		t.Fatal("failed to fetch header", err)
	}

	_, err = client.FetchVidCommonByHeight(ctx, blockHeight-1)
	if err != nil {
		t.Fatal("failed to fetch vid common", err)
	}

	_, err = client.FetchHeadersByRange(ctx, 1, 10)
	if err != nil {
		t.Fatal("failed to fetch headers by range", err)
	}

	_, err = client.FetchTransactionsInBlock(ctx, 1, 1)
	if err != nil {
		t.Fatal("failed to fetch transactions in block", err)
	}

	_, err = client.FetchBlockMerkleProof(ctx, 1, blockHeight)
	if err != nil {
		t.Fatal("failed to fetch block merkle proof", err)
	}
}

func getHeaderFromTestFile(path string, t *testing.T) types.HeaderInterface {
	file, err := os.Open(path)
	if err != nil {
		t.Fatal("failed to open file:", err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		t.Fatal("Error reading file:", err)
	}

	var headerImpl types.HeaderImpl
	err = json.Unmarshal(data, &headerImpl)
	if err != nil {
		t.Fatal("Error unmarshaling:", err)
	}

	return headerImpl.Header
}
