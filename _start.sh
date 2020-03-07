#!/bin/bash
go mod vendor
cd cmd/zeroslopesvc
swag init -parseDependency
go run .
cd ../..
