SHELL = /bin/bash -eu -o pipefail

export GOPATH?=$(shell go env GOPATH)
export CGO_ENABLED=0
export GOPROXY?=https://proxy.golang.org
export GO111MODULE=on
export GOFLAGS?=-mod=readonly -trimpath
export GIT_TAG ?= $(shell git tag --points-at HEAD)

GO_VERSION = 1.20.1
GO 				?= go
GOOS 	?= $(shell go env GOOS)
GOLANGCI_LINT := $(GO) run github.com/golangci/golangci-lint/cmd/golangci-lint@v1.52
GOFUMPT 		:= $(GO) run mvdan.cc/gofumpt@v0.4
GOARCH 	?= $(shell go env GOARCH)

CMD = $(notdir $(wildcard ./cmd/*))
BUILD_DEST ?= _build

REGISTRY ?= quay.io
REGISTRY_NAMESPACE ?= tinkerbell

IMAGE_TAG = \
		$(shell echo $$(git rev-parse HEAD && if [[ -n $$(git status --porcelain) ]]; then echo '-dirty'; fi)|tr -d ' ')
IMAGE_NAME ?= $(REGISTRY)/$(REGISTRY_NAMESPACE)/operator:$(IMAGE_TAG)

BASE64_ENC = \
		$(shell if base64 -w0 <(echo "") &> /dev/null; then echo "base64 -w0"; else echo "base64 -b0"; fi)

.PHONY: verify
verify: lint ## Verify code style, is lint free, freshness ...
	$(GOFUMPT) -d .

.PHONY: lint
lint: shellcheck hadolint golangci-lint yamllint ## Lint code

LINT_ARCH := $(shell uname -m)
LINT_OS := $(shell uname)
LINT_OS_LOWER := $(shell echo $(LINT_OS) | tr '[:upper:]' '[:lower:]')

SHELLCHECK_VERSION ?= v0.8.0
SHELLCHECK_BIN := out/linters/shellcheck-$(SHELLCHECK_VERSION)-$(LINT_ARCH)
$(SHELLCHECK_BIN):
	mkdir -p out/linters
	curl -sSfL -o $@.tar.xz https://github.com/koalaman/shellcheck/releases/download/$(SHELLCHECK_VERSION)/shellcheck-$(SHELLCHECK_VERSION).$(LINT_OS_LOWER).$(LINT_ARCH).tar.xz \
		|| echo "Unable to fetch shellcheck for $(LINT_OS)/$(LINT_ARCH): falling back to locally install"
	test -f $@.tar.xz \
		&& tar -C out/linters -xJf $@.tar.xz \
		&& mv out/linters/shellcheck-$(SHELLCHECK_VERSION)/shellcheck $@ \
		|| printf "#!/usr/bin/env shellcheck\n" > $@
	chmod u+x $@

.PHONY: shellcheck
shellcheck: $(SHELLCHECK_BIN)
	$(SHELLCHECK_BIN) $(shell find . -name "*.sh")

HADOLINT_VERSION ?= v2.12.1-beta
HADOLINT_BIN := out/linters/hadolint-$(HADOLINT_VERSION)-$(LINT_ARCH)
$(HADOLINT_BIN):
	mkdir -p out/linters
	curl -sSfL -o $@.dl https://github.com/hadolint/hadolint/releases/download/$(HADOLINT_VERSION)/hadolint-$(LINT_OS)-$(LINT_ARCH) \
		|| echo "Unable to fetch hadolint for $(LINT_OS)/$(LINT_ARCH), falling back to local install"
	test -f $@.dl && mv $(HADOLINT_BIN).dl $@ || printf "#!/usr/bin/env hadolint\n" > $@
	chmod u+x $@

.PHONY: hadolint
hadolint: $(HADOLINT_BIN)
	$(HADOLINT_BIN) --no-fail $(shell find . -name "*Dockerfile")

.PHONY: golangci-lint
golangci-lint:
	$(GOLANGCI_LINT) run

YAMLLINT_VERSION ?= 1.26.3
YAMLLINT_ROOT := out/linters/yamllint-$(YAMLLINT_VERSION)
YAMLLINT_BIN := $(YAMLLINT_ROOT)/dist/bin/yamllint
$(YAMLLINT_BIN):
	mkdir -p out/linters
	rm -rf out/linters/yamllint-*
	curl -sSfL https://github.com/adrienverge/yamllint/archive/refs/tags/v$(YAMLLINT_VERSION).tar.gz | tar -C out/linters -zxf -
	cd $(YAMLLINT_ROOT) && pip3 install --target dist . || pip install --target dist .

.PHONY: yamllint
yamllint: $(YAMLLINT_BIN)
	PYTHONPATH=$(YAMLLINT_ROOT)/dist $(YAMLLINT_ROOT)/dist/bin/yamllint .


.PHONY: vendor
vendor: buildenv
	go mod vendor

.PHONY: buildenv
buildenv:
	@go version

.PHONY: all
all: build

.PHONY: build
build:
	CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=$(GOARCH) $(GO) build $(LDFLAGS) -o ./bin/operator-$(GOOS)-$(GOARCH) ./cmd/tinkerbell

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
