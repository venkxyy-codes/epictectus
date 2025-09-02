# ====== Variables =======
PROJECT_NAME := "$(shell basename `git rev-parse --show-toplevel`)"
SSH_PRIVATE_KEY=`cat $(ssh_private_key)`
BUILD_DIR := "./out"
APP_EXECUTABLE="$(BUILD_DIR)/$(PROJECT_NAME)"

# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
ifeq (,$(shell go env GOBIN))
	GOBIN=$(shell go env GOPATH)/bin
else
	GOBIN=$(shell go env GOBIN)
endif

BUILD_INFO_GIT_TAG ?= $(shell git describe --tags 2>/dev/null || echo unknown)
BUILD_INFO_GIT_COMMIT ?= $(shell git rev-parse --short=10 HEAD 2>/dev/null || echo unknown)
BUILD_INFO_BUILD_DATE ?= $(shell date -u +"%Y-%m-%dT%H:%M:%SZ" || echo unknown)
BUILD_INFO_VERSION ?= $(shell prefix=$$(echo $(BUILD_INFO_GIT_TAG) | cut -c 1); if [ "$${prefix}" = "v" ]; then echo $(BUILD_INFO_GIT_TAG) | cut -c 2- ; else echo $(BUILD_INFO_GIT_TAG) ; fi)

GOOS := $(shell go env GOOS)
GOARCH := $(shell go env GOARCH)
GO_BUILD := GOOS=${GOOS} GOARCH=${GOARCH} go build
GO_RUN := go run $(LD_FLAGS)

# ====== Help =======
.PHONY: help
all: help
help: Makefile
	@echo
	@echo " Choose a command run in "$(PROJECT_NAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo

# ====== Targets =======

## start: Starts the go server
start: application.yaml
	@$(GO_RUN) main.go server --config-file config.yml

## build: Build a single binary
build:
	@echo "building executable"
	@rm -rf $(BUILD_DIR)
	@mkdir -p $(BUILD_DIR)
	@$(GO_BUILD) -o $(APP_EXECUTABLE) main.go

## test: Run all tests
test: .env test-run test-cov

test-run:
	@go test ./... -covermode=count -coverprofile=test.cov

config.yaml:
	@cp application.sample.yaml application.yaml
	@sed -i.bak 's/kafkacar/localhost/' application.yaml && rm application.yaml.bak

copy-config:
	cp config.example.yml config.yaml

lint: $(TOOLS_DIR)/golangci-lint
	$(TOOLS_DIR)/golangci-lint --max-issues-per-linter=0 --issues-exit-code=0 run

gomod.tidy:
	go mod tidy

pre-master-push: gomod.tidy lint test

tools-files: $(TOOLS_MOD_DIR)/go.mod $(TOOLS_MOD_DIR)/go.sum $(TOOLS_MOD_DIR)/tools.go

$(TOOLS_DIR)/gocov: tools-files
	@cd $(TOOLS_MOD_DIR) && \
	go build -tags=tools -o $(TOOLS_DIR)/gocov github.com/axw/gocov/gocov

$(TOOLS_DIR)/gci: tools-files
	@cd $(TOOLS_MOD_DIR) && \
	go build -tags=tools -o $(TOOLS_DIR)/gci github.com/daixiang0/gci

$(TOOLS_DIR)/golangci-lint: tools-files
	@cd $(TOOLS_MOD_DIR) && \
	go build -tags=tools -o $(TOOLS_DIR)/golangci-lint github.com/golangci/golangci-lint/cmd/golangci-lint

ci: build
	./$(APP_EXECUTABLE) migrate --stream="up" --config-file config.test.yml
