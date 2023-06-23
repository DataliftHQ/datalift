SHELL:=/usr/bin/env bash
.DEFAULT_GOAL:=all

MAKEFLAGS += --no-print-directory

DOCS_DEPLOY_USE_SSH ?= true
DOCS_DEPLOY_GIT_USER ?= git

# To update the version, update VERSION variable:
# Naming convention:
#   Stable releases:   "1.0.0"
#   Pre-releases:      "1.0.0-alpha.1", "1.0.0-beta.2", "1.0.0-rc.3"
#   Master/dev branch: "1.0.0-dev"
VERSION := 0.0.0-dev

YARN:=./build/bin/yarn.sh
PROJECT_ROOT_DIR:=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))

.PHONY: help # Print this help message.
 help:
	@grep -E '^\.PHONY: [a-zA-Z0-9_-]+ .*?# .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = "(: |#)"}; {printf "%-30s %s\n", $$2, $$3}'

.PHONY: proto # Generate protobuf assets.
proto:
	tools/buf.sh generate
	cd server && ../tools/buf.sh generate

.PHONY: proto-lint # Lint the generated protobuf assets.
proto-lint:
	tools/buf.sh lint
	cd server && ../tools/buf.sh lint

.PHONY: proto-verify # Verify proto changes include generate server assets.
proto-verify:
	find server/config -mindepth 1 -maxdepth 1 -type d -exec rm -rf {} \;
	$(MAKE) proto
	tools/ensure-no-diff.sh api server/config

.PHONY: server # Build the standalone server.
server: preflight-checks-server
	cd server && go build -o ../build/server -ldflags="-X main.version=$(VERSION)"

.PHONY: server-with-assets # Build the server with web assets.
server-with-assets: preflight-checks-server web
	cd server && go run cmd/assets/generate.go ../ui/build && go build -tags withAssets -o ../build/server -ldflags="-X main.version=$(VERSION)"

.PHONY: server-dev # Start the server in development mode.
server-dev: preflight-checks-server
	tools/air.sh

.PHONY: server-dev-mock # Start the server in development mode with mock responses.
server-dev-mock:
	cd server && go run mock/gateway.go

.PHONY: server-lint # Lint the server code.
server-lint: preflight-checks-server
	cd client && ../tools/golangci-lint.sh run --timeout 2m30s
	cd server && ../tools/golangci-lint.sh run --timeout 2m30s
	cd worker && ../tools/golangci-lint.sh run --timeout 2m30s

.PHONY: server-lint-fix # Lint and fix the server code.
server-lint-fix:
	tools/golangci-lint.sh run --fix
	cd server && go mod tidy

.PHONY: server-test # Run unit tests for the server code.
server-test: preflight-checks-server
	cd server && go test -race -covermode=atomic ./...

.PHONY: server-verify # Verify go modules' requirements files are clean.
server-verify:
	cd server && go mod tidy
	tools/ensure-no-diff.sh server

.PHONY: server-config-validation
server-config-validation:
	cd server && go run main.go -validate -c datalift-config.yaml

.PHONY: preflight-checks-server
preflight-checks-server:
	@tools/preflight-checks.sh server

.PHONY: web # Build production web assets.
web: yarn-ensure preflight-checks-web yarn-install
	$(YARN) --cwd ui build

.PHONY: web-dev-build # Build development web assets.
web-dev-build: yarn-install
	$(YARN) --cwd ui preview

.PHONY: web-dev # Start the web in development mode.
web-dev: yarn-install
	$(YARN) --cwd ui dev

.PHONY: web-lint # Lint the web code.
web-lint: yarn-ensure
	$(YARN) --cwd ui lint

.PHONY: web-lint-fix # Lint and fix the web code.
web-lint-fix: yarn-ensure
	$(YARN) --cwd ui lint:fix

.PHONY: web-test # Run unit tests for the web code.
web-test: yarn-ensure
	$(YARN) --cwd ui test

# TODO: FIX ME
.PHONY: web-verify # Verify web packages are sorted.
web-verify: yarn-ensure
	$(YARN) --cwd ui lint:packages

.PHONY: yarn-install # Install web dependencies.
yarn-install: yarn-ensure
	$(YARN) --cwd ui install --frozen-lockfile

.PHONY: yarn-ensure # Install the pinned version of yarn.
yarn-ensure:
	@tools/install-yarn.sh

.PHONY: preflight-checks-web
preflight-checks-web:
	@tools/preflight-checks.sh web

.PHONY: dev # Run the application in development mode.
dev:
	$(MAKE) -j2 server-dev web-dev

.PHONY: dev-mock # Run the application in development mode with mock responses.
dev-mock:
	$(MAKE) -j2 server-dev-mock web-dev

.PHONY: lint # Lint all of the code.
lint: proto-lint server-lint web-lint

.PHONY: lint-fix # Lint and fix all of the code.
lint-fix: server-lint-fix web-lint-fix

.PHONY: test # Unit test all of the code.
test: server-test web-test

.PHONY: verify # Verify all of the code.
verify: proto-verify server-verify web-verify

.PHONY: clean # Remove build and cache artifacts.
clean:
	rm -rf build
	cd ui && rm -rf build node_modules .yarn

.PHONY: dev-k8s-up # Start a local k8s cluster.
dev-k8s-up:
	@tools/kind.sh create cluster --kubeconfig $(PROJECT_ROOT_DIR)/build/kubeconfig-datalift --name datalift-local || true
	@tools/kind.sh seed

	@echo
	@echo "Export these environment variables before starting development:"
	@echo '    export KUBECONFIG=$(PROJECT_ROOT_DIR)/build/kubeconfig-datalift'

.PHONY: dev-k8s-down # Stop a local k8s cluster.
dev-k8s-down:
	@tools/kind.sh delete cluster --name datalift-local

.PHONY: preflight-checks
preflight-checks:
	@tools/preflight-checks.sh

.PHONY: preflight-checks-worker
preflight-checks-worker:
	@tools/preflight-checks.sh worker

.PHONY: worker # Build the standalone worker.
worker: preflight-checks-worker
	cd worker && go build -o ../build/worker -ldflags="-X main.version=$(VERSION)"
