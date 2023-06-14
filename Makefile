build-windows:
	env GOOS=windows GOARCH=amd64 go build -o eu4-checksum-patcher.exe *.go 