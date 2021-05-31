#!/bin/bash
cd cmd
GOOS=linux GOARCH=amd64 go build -o ../binaries/dekey-recover-linux ./dekey-recover.go
GOOS=darwin GOARCH=amd64 go build -o ../binaries/dekey-recover-mac ./dekey-recover.go
GOOS=windows GOARCH=amd64 go build -o ../binaries/dekey-recover-win.exe ./dekey-recover.go