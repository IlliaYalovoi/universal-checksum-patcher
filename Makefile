all: build-windows

build-windows:
	env GOOS=windows GOARCH=amd64 go build -o ./build/universal-checksum-patcher-windows.exe *.go