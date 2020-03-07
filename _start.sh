#!/bin/bash
dep ensure
cd cmd/zeroslopesvc
swag init -parseDependency
go run .
cd ../..
