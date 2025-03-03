package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"runtime"

	"github.com/spf13/cobra"
)

const targetDir = "../target"

const baseURL = "https://github.com/EspressoSystems/espresso-sequencer-go/releases"

func main() {
	var version string

	var rootCmd = &cobra.Command{Use: "app"}
	var downloadCmd = &cobra.Command{
		Use:   "download",
		Short: "Download the static library",
		Run: func(cmd *cobra.Command, args []string) {
			download(version)
		},
	}
	downloadCmd.Flags().StringVarP(&version, "version", "v", "latest", "Specify the version to download")

	var cleanCmd = &cobra.Command{
		Use:   "clean",
		Short: "Clean the downloaded files",
		Run: func(cmd *cobra.Command, args []string) {
			clean()
		},
	}

	rootCmd.AddCommand(downloadCmd, cleanCmd)
	err := rootCmd.Execute()
	if err != nil {
		fmt.Printf("Failed to execute command: %s\n", err)
		os.Exit(1)
	}
}

func download(version string) {
	fileName := getFileName()
	fileDir := getFileDir()
	libFilePath := filepath.Join(fileDir, fileName)

	if _, err := os.Stat(libFilePath); err == nil {
		fmt.Println("File already exists.")
		return
	}

	if err := os.MkdirAll(fileDir, 0755); err != nil {
		fmt.Printf("Failed to create target directory: %s\n", err)
		os.Exit(1)
	}

	url := fmt.Sprintf("%s/download/%s/%s", baseURL, version, fileName)

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

func clean() {
	fileDir := getFileDir()
	err := os.RemoveAll(fileDir)
	if err != nil {
		fmt.Printf("Failed to clean files: %s\n", err)
		os.Exit(1)
	}
	fmt.Println("Cleaned downloaded files.")
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
