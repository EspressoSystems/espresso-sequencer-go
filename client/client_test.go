package client

import (
	"context"
	"os/exec"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/log"
)

var workingDir = "./dev-node"

func TestApiWithEspressoDevNode(t *testing.T) {
	ctx := context.Background()
	cleanup := runEspresso(t, ctx)
	defer cleanup()

	err := waitForEspressoNode(t, ctx)
	if err != nil {
		t.Fatal("failed to start espresso dev node", err)
	}

	client := NewClient("http://localhost:21000")

	blockHeight, err := client.FetchLatestBlockHeight(ctx)
	if err != nil {
		t.Fatal("failed to fetch block height")
	}

	_, err = client.FetchHeaderByHeight(ctx, blockHeight)
	if err != nil {
		t.Fatal("failed to fetch header", err)
	}

	_, err = client.FetchVidCommonByHeight(ctx, blockHeight)
	if err != nil {
		t.Fatal("failed to fetch vid common", err)
	}

	_, err = client.FetchHeadersByRange(ctx, 1, blockHeight)
	if err != nil {
		t.Fatal("failed to fetch headers by range", err)
	}

}

func runEspresso(t *testing.T, ctx context.Context) func() {
	shutdown := func() {
		p := exec.Command("docker", "compose", "down")
		p.Dir = workingDir
		err := p.Run()
		if err != nil {
			panic(err)
		}
	}

	shutdown()
	invocation := []string{"compose", "up", "-d", "--build"}
	nodes := []string{
		"espresso-dev-node",
	}
	invocation = append(invocation, nodes...)
	procees := exec.Command("docker", invocation...)
	procees.Dir = workingDir

	go func() {
		if err := procees.Run(); err != nil {
			log.Error(err.Error())
			panic(err)
		}
	}()
	return shutdown
}

func waitForWith(
	t *testing.T,
	ctxinput context.Context,
	timeout time.Duration,
	interval time.Duration,
	condition func() bool,
) error {
	ctx, cancel := context.WithTimeout(ctxinput, timeout)
	defer cancel()

	for {
		if condition() {
			return nil
		}
		select {
		case <-time.After(interval):
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func waitForEspressoNode(t *testing.T, ctx context.Context) error {
	return waitForWith(t, ctx, 400*time.Second, 1*time.Second, func() bool {
		out, err := exec.Command("curl", "-s", "-L", "-f", "http://localhost:21000/availability/block/10").Output()
		if err != nil {
			log.Warn("error executing curl command:", "err", err)
			return false
		}

		return len(out) > 0
	})
}
