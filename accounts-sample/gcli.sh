#! /bin/bash

echo "pulling golangci-lint to install"
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.32.0

echo "GOPATH"
echo $GOPATH
echo "Directory"
echo $PWD

echo "Running golangci-lint -v run"
golangci-lint -v run --timeout 3m
