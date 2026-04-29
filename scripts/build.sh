#!/usr/bin/env bash
set -euo pipefail

VERSION="${VERSION:-dev}"

CGO_ENABLED=0 go build \
  -trimpath \
  -ldflags="-s -w -X main.version=${VERSION}" \
  -o ./bin/unc \
  ./cmd/unc
