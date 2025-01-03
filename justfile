lint:
    golangci-lint run ./...

test:
    go test -v ./...

bind-light-client:
	cd espresso-sequencer/contracts && forge build --force
	cd espresso-sequencer/contracts/out/LightClient.sol && cat LightClient.json | jq .abi > LightClient.abi
	cd espresso-sequencer/contracts/out/LightClientMock.sol && cat LightClientMock.json | jq .abi > LightClientMock.abi
	abigen --abi espresso-sequencer/contracts/out/LightClient.sol/LightClient.abi --pkg lightclient --out light-client/lightclient.go
	abigen --abi espresso-sequencer/contracts/out/LightClientMock.sol/LightClientMock.abi --pkg lightclientmock --out light-client-mock/lightclient.go

verification_dir := "./verification/rust"
target_lib := "./target/lib"

triple := if arch() == "aarch64" {
	if os() == "macos" {
		"aarch64-apple-darwin"
	} else {
		"aarch64-unknown-linux-gnu"
	}
} else if arch() == "x86_64" {
	if os() == "macos" {
		"x86_64-apple-darwin"
	} else {
		"x86_64-unknown-linux-gnu"
	}
} else {
	error("{{arch()}} is not supported")
}

build-verification:
	mkdir -p {{target_lib}}
	cargo build --release --manifest-path {{verification_dir}}/Cargo.toml --target {{triple}}
	install {{verification_dir}}/target/{{triple}}/release/libespresso_crypto_helper.a {{target_lib}}/libespresso_crypto_helper-{{triple}}.a
	go build ./verification
