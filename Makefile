# If the USE_SUDO_FOR_DOCKER env var is set, prefix docker commands with 'sudo'
ifdef USE_SUDO_FOR_DOCKER
	SUDO_CMD = sudo
endif

IMAGE ?= dinomiteX/go-test-server
TAG ?= 0.0.1
build: ## Builds the server
	GIT_TERMINAL_PROMPT=1 GOOS=linux GOARCH=amd64 CGO_ENABLED=0 GO111MODULE=on go build -o ./go-test-server --ldflags="-s" .

image: build ## Builds a Linux based image
	$(SUDO_CMD) docker build image/ -t "$(IMAGE):$(TAG)"

push: image ## Pushes the image to dockerhub, REQUIRES SPECIAL PERMISSION
	$(SUDO_CMD) docker push "$(IMAGE):$(TAG)"

.PHONY: build image push
