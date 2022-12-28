#!/bin/bash
# TODO: CI to generate coverage badge
# Generate coverage report and then generate badge

set -euo pipefail

go test -v ./... -covermode=count -coverprofile=coverage.out
go tool cover -func=coverage.out -o=coverage.out
RESULT=$(< coverage.out grep total | awk '{gsub(/%/,""); print $3}')
COLOR=GREEN # static color for now
curl "https://img.shields.io/badge/coverage-$RESULT%25-$COLOR" > badge.svg
