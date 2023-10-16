#!/bin/bash
set -eu

BINARY_NAME=mhnow-bot
BUILD_DIR=./build
declare -r BINARY_NAME
declare -r BUILD_DIR

if [[ ! -e "${BUILD_DIR}" ]]; then
    mkdir -p "${BUILD_DIR}"
fi

GOOS=linux   GOARCH=amd64 go build -o "${BUILD_DIR}/${BINARY_NAME}_linux_x86_64" ./main.go
GOOS=linux   GOARCH=arm64 go build -o "${BUILD_DIR}/${BINARY_NAME}_linux_arm64" ./main.go
GOOS=windows GOARCH=amd64 go build -o "${BUILD_DIR}/${BINARY_NAME}_windows_x86_64.exe" ./main.go
GOOS=darwin  GOARCH=arm64 go build -o "${BUILD_DIR}/${BINARY_NAME}_macos_arm64" ./main.go
