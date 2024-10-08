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
