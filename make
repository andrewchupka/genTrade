#!/bin/bash
echo "Building"
go build -o build.out
echo "Build successful. Running main.go"
go run main.go > output.txt
less output.txt