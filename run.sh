#!/usr/bin/env bash
set -euo pipefail

# Expects configuration from .env (sourced manually).
#
# Usage:
#   source .env && ./run.sh

PROJECT_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$PROJECT_ROOT"

if ! command -v go >/dev/null 2>&1; then
  echo "[run.sh] Go is not installed or not in PATH." >&2
  exit 1
fi

echo "[run.sh] Ensuring Go deps..."
go mod tidy

echo "[run.sh] Building..."
go build -o ./bin/server ./main.go

echo "[run.sh] Starting server on ${HTTP_LISTEN_ADDRESS:-:9000}"
exec ./bin/server
