#!/bin/bash
glide install
cd cmd/zeroslopesvc
swag init 
go run .
cd ../..
