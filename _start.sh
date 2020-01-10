#!/bin/bash
glide install
cd cmd/zeroslopesvc
swag init -parseDependency
go run .
cd ../..
