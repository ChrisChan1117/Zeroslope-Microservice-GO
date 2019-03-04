#!/bin/bash
glide install
cd src; swag init; cd ..
go run src/main.go