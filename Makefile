NAME = picsorter
VERSION = $(shell cat VERSION)
BUILD_GIT_SHA = $(shell git log -n 1 --pretty=format:"%H")
GIT_REMOTE = $(shell git remote get-url --push origin)
SSH_USERNAME ?= $(USER)
IMAGE = creisor/$(NAME)
IMAGE_TAGGED = $(IMAGE):$(BUILD_GIT_SHA)

.PHONY: help
help: ## Shows the help
	@echo 'USAGE: make <TARGETS>'
	@echo ''
	@echo 'Available targets are:'
	@echo ''
	@grep -E '^[ a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
	awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

all: help

.PHONY: vendor
vendor: ## Vendorize the go deps
	go mod tidy
	go mod vendor

.PHONY: test
test: unit lint ## Runs the unit tests and linters

.PHONY: install-build-deps
install-build-deps: ## Install the tools needed for testing, linting, and building
	go get github.com/gordonklaus/ineffassign
	go get golang.org/x/lint/golint
	go get github.com/axw/gocov/gocov
	go get github.com/AlekSi/gocov-xml
	go get github.com/securego/gosec/cmd/gosec

.PHONY: lint
lint: ## Run linter and static analysis
	@echo $@
	golint $$(go list ./...)
	ineffassign .
	gosec -quiet ./...

.PHONY: unit
unit: ## Runs the unit tests
	go test ./...

.PHONY: go-build
go-build: ## builds the binary
	go build -ldflags "-X main.Version=$(VERSION) -X main.Build=$(BUILD_GIT_SHA)" -o $(NAME)

.PHONY: go-build-all
go-build-all: go-build-mac go-build-linux ## builds the binaries for all supported platforms

.PHONY: go-build-mac
go-build-mac: ## builds the binary for MacOS
	GOOS=darwin GOARCH=amd64 go build -ldflags "-X main.Version=$(VERSION) -X main.Build=$(BUILD_GIT_SHA)" -o build/$(NAME)-darwin-amd64

.PHONY: go-build-linux
go-build-linux: ## builds the binary for MacOS
	GOOS=linux GOARCH=amd64 go build -ldflags "-X main.Version=$(VERSION) -X main.Build=$(BUILD_GIT_SHA)" -o build/$(NAME)-linux-amd64

.PHONY: docker-build-image
docker-build-image: ## Build the docker image
	docker build --rm -t $(IMAGE) -t $(IMAGE_TAGGED) \
	--build-arg BUILD_GIT_SHA=$(BUILD_GIT_SHA) \
	--build-arg NAME=$(NAME) \
	--build-arg BUILT_BY=$(SSH_USERNAME) \
	.
	#--build-arg GIT_REMOTE=$(GIT_REMOTE) \

.PHONY: docker-run
docker-run: ## Run the container
	@docker run -v ${PWD}/IMG_0975.jpg:/app/IMG_0975.jpg -it --rm $(IMAGE)
