#!/bin/sh

# Cleaning the Go cache only makes sense if we actually have Go installed... or
# if Go is actually callable. This does not hold true during deb packaging, so
# we need an explicit check to avoid build failures.
if ! command -v go > /dev/null; then
  exit
fi

DIR="bin"

if [ -d "$DIR" ]; then
  echo "Building source in ${DIR}"
else
  echo "Directory: ${DIR} not found. Creating"
  echo "Building source in ${DIR}"
fi

go mod download
go build -o $DIR/test cmd/test/main.go