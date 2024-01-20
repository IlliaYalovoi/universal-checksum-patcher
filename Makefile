all: build-windows
# Not powershell compatible. Meant to be used from unix shell
build-windows:
	env GOOS=windows GOARCH=amd64 go build -o ./build/universal-checksum-patcher.exe *.go