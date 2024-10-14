package types

import (
	"encoding/json"
	"io"
	"os"
	"testing"
)

func TestVersion(t *testing.T) {
	s := `{"Version":{"major":0,"minor":3}}`
	data := []byte(s)
	var v Version
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
