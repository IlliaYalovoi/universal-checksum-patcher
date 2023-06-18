all: build-windows build-linux

build-windows:
	env GOOS=windows GOARCH=amd64 go build -o ./build/eu4-checksum-patcher-windows.exe *.go 

build-linux:
	env GOOS=linux GOARCH=amd64 go build -o ./build/eu4-checksum-patcher-linux *.go 