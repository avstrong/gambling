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

COLOR := "\e[1;36m%s\e[0m\n"

.PHONY: register-user
register-user: ## Register a new user with dynamic email
	@if [ -z "$(email)" ]; then \
		echo "ERROR: 'email' is required"; \
		exit 1; \
	fi
	@printf $(COLOR) "Registering $(email) ..."
	@curl -X POST http://localhost:9092/api/v1/user \
	-H "Content-Type: application/json" \
	-d '{"email": "$(email)"}'

.PHONY: deposit-funds
deposit-funds: ## Deposit funds for a user with dynamic amount and currency, defaults to 10 USD
	@if [ -z "$(user)" ]; then \
		echo "ERROR: 'user' is required"; \
		exit 1; \
	fi
	$(eval amount := $(if $(amount),$(amount),10))
	$(eval currency := $(if $(currency),$(currency),USD))
	@printf $(COLOR) "Depositing $(amount) $(currency) for $(user) ..."
	@curl -X POST http://localhost:9092/api/v1/wallet/deposit \
	-H "Content-Type: application/json" \
	-d '{"userID": "$(user)", "amount": $(amount), "currency": "$(currency)"}'

.PHONY: withdraw-funds
withdraw-funds: ## Withdraw funds for a user with dynamic amount and currency, defaults to 10 USD
	@if [ -z "$(user)" ]; then \
		echo "ERROR: 'user' is required"; \
		exit 1; \
	fi
	$(eval amount := $(if $(amount),$(amount),10))
	$(eval currency := $(if $(currency),$(currency),USD))
	@printf $(COLOR) "Withdrawing $(amount) $(currency) for $(user) ..."
	@curl -X POST http://localhost:9092/api/v1/wallet/withdraw \
	-H "Content-Type: application/json" \
	-d '{"userID": "$(user)", "amount": $(amount), "currency": "$(currency)"}'

.PHONY: check-balance
check-balance: ## Check balance of a user, defaults to USD
	@if [ -z "$(user)" ]; then \
		echo "ERROR: 'user' is required"; \
		exit 1; \
	fi
	$(eval currency := $(if $(currency),$(currency),USD))
	@printf $(COLOR) "Checking balance for $(user) in $(currency) ..."
	@curl -X GET http://localhost:9092/api/v1/wallet/balance \
	-H "Content-Type: application/json" \
	-d '{"userID": "$(user)", "currency": "$(currency)"}'

.PHONY: create-game
create-game: ## Create a new game with optional attributes, defaults to 2 maxPlayers, 10 entryFee, and USD
	$(eval maxPlayers := $(if $(maxPlayers),$(maxPlayers),2))
	$(eval entryFee := $(if $(entryFee),$(entryFee),10))
	$(eval entryCurrency := $(if $(entryCurrency),$(entryCurrency),USD))
	@printf $(COLOR) "Creating a new game ..."
	@curl -X POST http://localhost:9092/api/v1/game \
	-H "Content-Type: application/json" \
	-d '{"name": "game 1", "maxPlayers": $(maxPlayers), "entryFee": $(entryFee), "entryCurrency": "$(entryCurrency)"}'

.PHONY: register-player
register-player: ## Register a player for a game with dynamic choice, defaults to false
	@if [ -z "$(gameID)" ] || [ -z "$(user)" ]; then \
		echo "ERROR: 'gameID' and 'user' are required"; \
		exit 1; \
	fi
	$(eval playerChoice := $(if $(playerChoice),$(playerChoice),false))
	@printf $(COLOR) "Registering a player for gameID $(gameID) with choice $(playerChoice) ..."
	@curl -X POST http://localhost:9092/api/v1/game/register \
	-H "Content-Type: application/json" \
	-d '{"gameID": "$(gameID)", "userID": "$(user)", "playerChoice": $(playerChoice)}'

.PHONY: play-game
play-game: ## Start the game with dynamic gameID
	@if [ -z "$(gameID)" ]; then \
		echo "ERROR: 'gameID' is required"; \
		exit 1; \
	fi
	@printf $(COLOR) "Starting the game with gameID $(gameID) ..."
	@curl -X POST http://localhost:9092/api/v1/game/play \
	-H "Content-Type: application/json" \
	-d '{"id": "$(gameID)"}'
