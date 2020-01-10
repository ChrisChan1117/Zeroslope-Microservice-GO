#!/bin/bash
glide install
cd cmd/zeroslopesvc
swag init -parseDependency
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .
cd ../..
