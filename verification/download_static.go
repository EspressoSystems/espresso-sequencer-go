package verification

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
)

const targetDir = "../target/lib"

const baseURL = "https://github.com/EspressoSystems/espresso-sequencer-go/releases/latest/download"

func init() {
	fileName := getFileName()
	libFilePath := filepath.Join(targetDir, fileName)

	if _, err := os.Stat(libFilePath); err == nil {
		return
	}

	if err := os.MkdirAll(targetDir, 0755); err != nil {
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
		fileName = "x86_64-unknown-linux-gnu"
	case "arm64":
		fileName = "aarch64-unknown-linux-gnu"
	default:
		panic(fmt.Sprintf("unsupported: %s", arch))
	}

	return fmt.Sprintf("libespresso_crypto_helper-%s.a", fileName)
}
