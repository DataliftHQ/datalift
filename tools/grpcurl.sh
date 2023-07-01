#!/usr/bin/env bash

REPO_ROOT="$(realpath "$(dirname "${BASH_SOURCE[0]}")/..")"
export GOBIN="${REPO_ROOT}/build/bin"

function main() {
  install
  "${GOBIN}/grpcurl" -plaintext "$@"
}

function install() {
  go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
}

main "$@"
