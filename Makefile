all: build-windows build-linux build-macos-amd64 build-macos-arm64

build-windows:
	env GOOS=windows GOARCH=amd64 go build -o ./build/universal-checksum-patcher-windows.exe *.go 

build-linux:
	env GOOS=linux GOARCH=amd64 go build -o ./build/universal-checksum-patcher-linux *.go 

build-macos-amd64:
	env GOOS=darwin GOARCH=amd64 go build -o ./build/universal-checksum-patcher-macos-amd64 *.go

build-macos-arm64:
	env GOOS=darwin GOARCH=arm64 go build -o ./build/universal-checksum-patcher-macos-arm64 *.go 