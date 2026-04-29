#!/usr/bin/env bash
set -euo pipefail

echo "Running tests..."
go test -json ./... | go run github.com/mfridman/tparse@latest -all
