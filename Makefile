NAME = picsorter
BUILD_GIT_SHA = $(shell git log -n 1 --pretty=format:"%h")
GIT_REMOTE = $(shell git remote get-url --push origin)
SSH_USERNAME ?= $(USER)
#IMAGE = creisor/$(NAME)
#IMAGE_TAGGED = $(IMAGE):$(BUILD_GIT_SHA)
IMAGE = c295f68afc30

.PHONY: help
help: ## Shows the help
	@echo 'USAGE: make <TARGETS>'
	@echo ''
	@echo 'Available targets are:'
	@echo ''
	@grep -E '^[ a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
	awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

all: help

.PHONY: test
test: unit ## Runs the unit tests

.PHONY: unit
unit: ## Runs the unit tests
	@docker run -it --rm go test ./...

.PHONY: docker-build-image
docker-build-image: ## Build the docker image
	docker build --rm -t $(IMAGE) -t $(IMAGE_TAGGED) \
	--build-arg BUILD_GIT_SHA=$(BUILD_GIT_SHA) \
	--build-arg NAME=$(NAME) \
	--build-arg BUILT_BY=$(SSH_USERNAME) \
	--build-arg GIT_REMOTE=$(GIT_REMOTE) \

.PHONY: docker-run
docker-run: ## Run the container
	@docker run -v ${PWD}/IMG_0975.jpg:/app/IMG_0975.jpg -it --rm $(IMAGE)
