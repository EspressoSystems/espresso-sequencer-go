package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"runtime"
)

const targetDir = "../target/lib"

const baseURL = "https://github.com/ImJeremyHe/espresso-sequencer-go/releases/latest/download"

func main() {
	fileName := getFileName()
	fileDir := getFileDir()
	libFilePath := filepath.Join(fileDir, fileName)

	if _, err := os.Stat(libFilePath); err == nil {
		return
	}

	if err := os.MkdirAll(fileDir, 0755); err != nil {
		fmt.Printf("Failed to create target directory: %s\n", err)
		os.Exit(1)
	}

	url := fmt.Sprintf("%s/%s", baseURL, fileName)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Failed to download static library: %s\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	out, err := os.Create(libFilePath)
	if err != nil {
		fmt.Printf("Failed to create file: %s\n", err)
		os.Exit(1)
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		fmt.Printf("Failed to write file: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Static library downloaded to: %s\n", libFilePath)
}

func getFileName() string {
	arch := runtime.GOARCH

	var fileName string

	switch arch {
	case "amd64":
		if runtime.GOOS == "darwin" {
			fileName = "x86_64-apple-darwin"
		} else if runtime.GOOS == "linux" {
			fileName = "x86_64-unknown-linux-musl"
		}
	case "arm64":
		if runtime.GOOS == "darwin" {
			fileName = "aarch64-apple-darwin"
		} else if runtime.GOOS == "linux" {
			fileName = "aarch64-unknown-linux-musl"
		}
	default:
		panic(fmt.Sprintf("unsupported: %s", arch))
	}

	return fmt.Sprintf("libespresso_crypto_helper-%s.a", fileName)
}

func getFileDir() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}

	return filepath.Join(path.Dir(filename), targetDir)
}
