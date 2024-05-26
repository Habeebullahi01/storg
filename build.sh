#!/bin/bash

APP_NAME="storg"
PLATFORMS=("windows/amd64" "darwin/amd64" "linux/amd64")

mkdir -p builds

for PLATFORM in "${PLATFORMS[@]}"
do
    OS=$(echo $PLATFORM | cut -d'/' -f1)
    ARCH=$(echo $PLATFORM | cut -d'/' -f2)
    OUTPUT="builds/$APP_NAME-$OS-$ARCH"
    if [ $OS = "windows" ]; then
        OUTPUT+=".exe"
    fi

    GOOS=$OS GOARCH=$ARCH go build -o $OUTPUT main.go
    if [ $? -ne 0 ]; then
        echo "Failed to build for $PLATFORM"
        exit 1
    fi
done
