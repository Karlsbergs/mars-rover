#!/bin/sh
echo "Run tests ..."

go test  -v -coverprofile=profile.cov ./...
go tool cover -func profile.cov

echo 
echo 
