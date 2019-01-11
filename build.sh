#!/bin/sh

GO111MODULE="on"

go mod download
go get
go vet
go build -o getpod .
