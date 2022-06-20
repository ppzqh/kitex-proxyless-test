#!/bin/bash

# build
go mod tidy
go build -o output/bin/kitex
cp ./bootstrap.sh output/