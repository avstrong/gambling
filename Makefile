ifeq ($(OS),Windows_NT)
    GOOS := windows
else
    UNAME_S := $(shell uname -s)
    ifeq ($(UNAME_S),Linux)
        GOOS := linux
    endif
    ifeq ($(UNAME_S),Darwin)
        GOOS := darwin
    endif
endif

PROJECT_PATH := $(shell pwd)
GO = $(shell which go)
SHA_COMMIT = $(shell git rev-parse --short HEAD)
GOARCH = amd64
CGO_ENABLED = 0
GO_BUILD = GOOS=$(GOOS) GOARCH=$(GOARCH) CGO_ENABLED=$(CGO_ENABLED) $(GO) build

COLOR := "\e[1;36m%s\e[0m\n"

.PHONY: help
help: ## show this help
	@egrep -h '\s##\s' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

.PHONY: build-app
build-app: ## Build the Go app
	@printf $(COLOR) "Build app ..."
	${GO_BUILD} -ldflags="-X 'main.shaSHA=$(SHA_COMMIT)' -X 'main.buildUser=$(id -u -n)' -X 'main.buildTime=$(date)'" -mod vendor -o ./bin/app

.PHONY: prepare-to-commit
prepare-to-commit: ## Set up the prepare-commit-msg hook for automatically including branch name in commit messages
	@printf $(COLOR) "Init local prepare-commit-msg ..."
	@cp -v ${PROJECT_PATH}/scripts/prepare-commit-msg ${PROJECT_PATH}/.git/hooks/prepare-commit-msg
	@printf $(COLOR) "Give permissions to .git/hooks/prepare-commit-msg..."
	@chmod +x .git/hooks/prepare-commit-msg

.PHONY: init-local-env
init-local-env: prepare-to-commit ## init local environment
	@printf $(COLOR) "Init .env file ..."
	@cp -v ${PROJECT_PATH}/.env.example ${PROJECT_PATH}/.env

.PHONY: test
test: ## run app tests
	@printf $(COLOR) "Run tests ..."
	$(GO) generate ./...
	$(GO) test -race -cover -short -count=1 \
  				-coverprofile profile.cov.tmp \
  				-coverpkg=./... \
  				./...
	cat profile.cov.tmp | grep -Ev "_gen.go|mock_" > profile.cov
	make cover

.PHONY: cover
cover:
	@printf $(COLOR) "Code coverage ..."
	$(GO) tool cover -func profile.cov

.PHONY: lint
lint: ## run linters
	@printf $(COLOR) "Run golangci-lint ..."
	@VERSION="1.52.2"; \
	CURRENT_VERSION=$$(golangci-lint --version 2> /dev/null | grep -Eo 'v[0-9]+\.[0-9]+\.[0-9]+'); \
	if ! which golangci-lint > /dev/null || [ "$$VERSION" != "$$CURRENT_VERSION" ]; then \
		${GO} install github.com/golangci/golangci-lint/cmd/golangci-lint@v$$VERSION; \
	fi
	@golangci-lint run ./... --color always
	@printf $(COLOR) "Run linters ..."

.PHONY: generate-proto
generate-proto:
	PATH="${PATH}:${HOME}/go/bin" protoc -I . \
    	--go_out=pkg/game --go_opt=paths=import \
        --go-grpc_out=pkg/game --go-grpc_opt=paths=import \
        --grpc-gateway_out=pkg/game \
        --grpc-gateway_opt=logtostderr=true \
        --grpc-gateway_opt=paths=import \
        api/proto/server.proto
