#!/usr/bin/env bash
set -euo pipefail

REPO_ROOT="${REPO_ROOT:-"$(realpath "$(dirname "${BASH_SOURCE[0]}")/..")"}"
BUILD_ROOT="${REPO_ROOT}/build"
BUILD_BIN="${BUILD_ROOT}/bin"

NAME=air
RELEASE=v1.44.0
OSX_RELEASE_SUM=f32f7fcc19fedf335f07e8757d06251375f977d63ea26b043a85ee2039b36588
LINUX_RELEASE_SUM=148e6b101d3a102674aff86faff36423597e3437d7da1dd8edac3bc27d7b8e49

ARCH=amd64

RELEASE_BINARY="${BUILD_BIN}/${NAME}-${RELEASE}"

ensure_binary() {
  if [[ ! -f "${RELEASE_BINARY}" ]]; then
    echo "info: Downloading ${NAME} ${RELEASE} to build environment"
    mkdir -p "${BUILD_BIN}"

    case "${OSTYPE}" in
      "darwin"*) os_type="darwin"; sum="${OSX_RELEASE_SUM}" ;;
      "linux"*) os_type="linux"; sum="${LINUX_RELEASE_SUM}" ;;
      *) echo "error: Unsupported OS '${OSTYPE}' for shellcheck install, please install manually" && exit 1 ;;
    esac

    release_archive="/tmp/${NAME}-${RELEASE}"
    URL="https://github.com/cosmtrek/air/releases/download/${RELEASE}/air_${RELEASE:1}_${os_type}_${ARCH}"
    curl -sSL -o "${release_archive}" "${URL}"
    echo ${sum} ${release_archive} | sha256sum --check --quiet -

    find "${BUILD_BIN}" -maxdepth 1 -regex '.*/'${NAME}'-[A-Za-z0-9\.]+$' -exec rm {} \;  # cleanup older versions
    mv "${release_archive}" "${RELEASE_BINARY}"
    chmod +x "${RELEASE_BINARY}"
  fi
}

ensure_fd() {
  if [[ "${OSTYPE}" == *"darwin"* ]]; then
    ulimit -n 1024
  fi
}

ensure_binary
ensure_fd

"${RELEASE_BINARY}"
