SHELL:=/usr/bin/env bash
.DEFAULT_GOAL:=all

MAKEFLAGS += --no-print-directory

DOCS_DEPLOY_USE_SSH ?= true
DOCS_DEPLOY_GIT_USER ?= git

VERSION := 0.0.0

YARN:=./build/bin/yarn.sh
PROJECT_ROOT_DIR:=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))

.PHONY: help # Print this help message.
 help:
	@grep -E '^\.PHONY: [a-zA-Z0-9_-]+ .*?# .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = "(: |#)"}; {printf "%-30s %s\n", $$2, $$3}'

.PHONY: proto # Generate protobuf assets.
proto:
	buf generate
	cd server && buf generate

.PHONY: proto-lint # Lint the generated protobuf assets.
proto-lint:
	buf lint
	cd server && buf lint

.PHONY: proto-verify # Verify proto changes include generate server assets.
proto-verify:
	find server/config -mindepth 1 -maxdepth 1 -type d -exec rm -rf {} \;
	$(MAKE) proto
	tools/ensure-no-diff.sh server/api server/config

.PHONY: server # Build the standalone server.
server: preflight-checks-server
	cd server && go build -o ../build/server -ldflags="-X main.version=$(VERSION)"

.PHONY: server-with-assets # Build the server with frontend assets.
server-with-assets: preflight-checks-server
	cd server && go run cmd/assets/generate.go ../frontend/build && go build -tags withAssets -o ../build/server -ldflags="-X main.version=$(VERSION)"

.PHONY: server-dev # Start the server in development mode.
server-dev: preflight-checks-server
	tools/air.sh

.PHONY: server-dev-mock # Start the server in development mode with mock responses.
server-dev-mock:
	cd server && go run mock/gateway.go

.PHONY: server-lint # Lint the server code.
server-lint: preflight-checks-server
	tools/golangci-lint.sh run --timeout 2m30s

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

.PHONY: frontend # Build production frontend assets.
frontend: yarn-ensure preflight-checks-frontend yarn-install
	$(YARN) --cwd frontend build

.PHONY: frontend-dev-build # Build development frontend assets.
frontend-dev-build: yarn-install
	$(YARN) --cwd frontend preview

.PHONY: frontend-dev # Start the frontend in development mode.
frontend-dev: yarn-install
	$(YARN) --cwd frontend dev

.PHONY: frontend-lint # Lint the frontend code.
frontend-lint: yarn-ensure
	$(YARN) --cwd frontend lint

.PHONY: frontend-lint-fix # Lint and fix the frontend code.
frontend-lint-fix: yarn-ensure
	$(YARN) --cwd frontend lint:fix

.PHONY: frontend-test # Run unit tests for the frontend code.
frontend-test: yarn-ensure
	$(YARN) --cwd frontend test

# TODO: FIX ME
.PHONY: frontend-verify # Verify frontend packages are sorted.
frontend-verify: yarn-ensure
	$(YARN) --cwd frontend lint:packages

.PHONY: yarn-install # Install frontend dependencies.
yarn-install: yarn-ensure
	$(YARN) --cwd frontend install --frozen-lockfile

.PHONY: yarn-ensure # Install the pinned version of yarn.
yarn-ensure:
	@./tools/install-yarn.sh

.PHONY: preflight-checks-frontend
preflight-checks-frontend:
	@tools/preflight-checks.sh frontend

.PHONY: dev # Run the Clutch application in development mode.
dev:
	$(MAKE) -j2 server-dev frontend-dev

.PHONY: dev-mock # Run the Clutch application in development mode with mock responses.
dev-mock:
	$(MAKE) -j2 server-dev-mock frontend-dev

.PHONY: lint # Lint all of the code.
lint: proto-lint server-lint frontend-lint

.PHONY: lint-fix # Lint and fix all of the code.
lint-fix: server-lint-fix frontend-lint-fix

.PHONY: test # Unit test all of the code.
test: server-test frontend-test

.PHONY: verify # Verify all of the code.
verify: proto-verify server-verify frontend-verify

.PHONY: clean # Remove build and cache artifacts.
clean:
	rm -rf build frontend/build frontend/node_modules frontend/.yarn

.PHONY: dev-k8s-up # Start a local k8s cluster.
dev-k8s-up:
	@tools/kind.sh create cluster --kubeconfig $(PROJECT_ROOT_DIR)/build/kubeconfig-clutch --name datalift-local || true
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
