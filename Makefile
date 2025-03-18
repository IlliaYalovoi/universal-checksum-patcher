all: go-test build-ps

build-bash:
	@env GOOS=windows GOARCH=amd64 go build -o ./build/universal-checksum-patcher.exe *.go
	@echo Successfully built universal-checksum-patcher.exe

build-ps:
	@powershell -Command "$$env:GOOS='windows'; $$env:GOARCH='amd64'; go build -o .\\build\\universal-checksum-patcher.exe ."
	@echo Successfully built universal-checksum-patcher.exe

go-test:
	@go test -v ./...