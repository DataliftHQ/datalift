#!/usr/bin/env bash
set -euo pipefail

REPO_ROOT="$(realpath "$(dirname "${BASH_SOURCE[0]}")/..")"
BUILD_ROOT="${REPO_ROOT}/build"
BUILD_BIN="${BUILD_ROOT}/bin"

NAME=buf
RELEASE=v1.21.0
OSX_RELEASE_256=7c1c0c2d045ce3aee1db3450014a7d8b978acae38a643d9319233c81c0f065df
LINUX_RELEASE_256=1db51318e49f12095c97866c9b5d939dfec318b50362bba8a3a9545c4cff456b

ARCH=x86_64

RELEASE_BINARY="${BUILD_BIN}/${NAME}-${RELEASE}"

main() {
  ensure_binary

  "${RELEASE_BINARY}" "$@"
}

ensure_binary() {
  if [[ ! -f "${RELEASE_BINARY}" ]]; then
    echo "info: Downloading ${NAME} ${RELEASE} to build environment"
    mkdir -p "${BUILD_BIN}"

    case "${OSTYPE}" in
      "darwin"*) os_type="Darwin"; sum="${OSX_RELEASE_256}" ;;
      "linux"*) os_type="Linux"; sum="${LINUX_RELEASE_256}" ;;
      *) echo "error: Unsupported OS '${OSTYPE}' for ${NAME} install, please install manually" && exit 1 ;;
    esac

    release_archive="/tmp/${NAME}-${RELEASE}.tar.gz"
    curl -sSL -o "${release_archive}" \
      "https://github.com/bufbuild/buf/releases/download/${RELEASE}/buf-${os_type}-${ARCH}.tar.gz"
    echo ${sum} ${release_archive} | sha256sum --check --quiet -

    release_tmp_dir="/tmp/${NAME}-${RELEASE}"
    mkdir -p "${release_tmp_dir}"
    tar -xzf "${release_archive}" --strip=1 -C "${release_tmp_dir}"

    if [[ ! -f "${RELEASE_BINARY}" ]]; then
      find "${BUILD_BIN}" -maxdepth 0 -regex '.*/'${NAME}'-[A-Za-z0-9\.]+$' -exec rm {} \;  # cleanup older versions
      mv "${release_tmp_dir}/bin/${NAME}" "${RELEASE_BINARY}"
    fi

    # Cleanup stale resources.
    rm "${release_archive}"
    rm -rf "${release_tmp_dir}"
  fi
}

main "$@"