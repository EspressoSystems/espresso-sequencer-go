default:
	just --list

lint:
    golangci-lint run ./...

test:
    go test -v ./...

bind:
	cd espresso-sequencer/contracts && forge build --force
	cd espresso-sequencer/contracts/out/HotShot.sol && cat HotShot.json | jq .abi > HotShot.abi
	abigen --abi espresso-sequencer/contracts/out/HotShot.sol/HotShot.abi --pkg hotshot --out hotshot/hotshot.go
