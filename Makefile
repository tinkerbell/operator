SHELL = /bin/bash -eu -o pipefail

export GOPATH?=$(shell go env GOPATH)
export CGO_ENABLED=0
export GOPROXY?=https://proxy.golang.org
export GO111MODULE=on
export GOFLAGS?=-mod=readonly -trimpath
export GIT_TAG ?= $(shell git tag --points-at HEAD)

GO_VERSION = 1.20.1

CMD = $(notdir $(wildcard ./cmd/*))
BUILD_DEST ?= _build

REGISTRY ?= docker.io
REGISTRY_NAMESPACE ?= moadqassem

IMAGE_TAG = \
		$(shell echo $$(git rev-parse HEAD && if [[ -n $$(git status --porcelain) ]]; then echo '-dirty'; fi)|tr -d ' ')
IMAGE_NAME ?= $(REGISTRY)/$(REGISTRY_NAMESPACE)/kubetink:$(IMAGE_TAG)

BASE64_ENC = \
		$(shell if base64 -w0 <(echo "") &> /dev/null; then echo "base64 -w0"; else echo "base64 -b0"; fi)

.PHONY: lint
lint:
	@golangci-lint --version
	golangci-lint run -v ./pkg/...

.PHONY: vendor
vendor: buildenv
	go mod vendor

.PHONY: buildenv
buildenv:
	@go version

.PHONY: all
all: build

.PHONY: build
build: $(CMD)

.PHONY: $(CMD)
$(CMD): %: $(BUILD_DEST)/%

$(BUILD_DEST)/%: cmd/%
	go build -v -o $@ ./cmd/$*

.PHONY: clean
clean:
	rm -rf $(BUILD_DEST)
	@echo "Cleaned $(BUILD_DEST)"

.PHONY: docker-image
docker-image:
	docker build --build-arg GO_VERSION=$(GO_VERSION) -t $(IMAGE_NAME) .

.PHONY: docker-image-publish
docker-image-publish: docker-image
	docker push $(IMAGE_NAME)
	if [[ -n "$(GIT_TAG)" ]]; then \
		$(eval IMAGE_TAG = $(GIT_TAG)) \
		docker build -t $(IMAGE_NAME) . && \
		docker push $(IMAGE_NAME) && \
		$(eval IMAGE_TAG = latest) \
		docker build -t $(IMAGE_NAME) . ;\
		docker push $(IMAGE_NAME) ;\
	fi

.PHONY: shfmt
shfmt:
	shfmt -w -sr -i 2 hack
