#!/usr/bin/env bash

set -e

if [ "$#" -lt 1 ]; then
  echo "Usage: $0 <MARKDOWN_PATH>..."
  exit 1
fi

go run ./internal/cmd/docextractor embed "$@"
