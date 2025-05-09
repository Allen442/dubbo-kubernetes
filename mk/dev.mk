DUBBO_DIR ?= .
TOOLS_DIR = $(DUBBO_DIR)/tools

CI_TOOLS_VERSION ?= master
CHART_REPO_NAME ?= dubbo
PROJECT_NAME ?= dubbo

CI_TOOLS_DIR ?= ${HOME}/.dubbo-dev/${PROJECT_NAME}-${CI_TOOLS_VERSION}
ifdef XDG_DATA_HOME
	CI_TOOLS_DIR := ${XDG_DATA_HOME}/dubbo-dev/${PROJECT_NAME}-${CI_TOOLS_VERSION}
endif
CI_TOOLS_BIN_DIR=$(CI_TOOLS_DIR)/bin

# Change here and `make check` ensures these are used for CI
# Note: These are _docker image tags_
# If changing min version, update mk/kind.mk as well
K8S_MIN_VERSION = v1.23.17-k3s1
K8S_MAX_VERSION = v1.28.1-k3s1
export GO_VERSION=$(shell go mod edit -json | jq -r .Go)
export GOLANGCI_LINT_VERSION=v1.60.2
GOOS := $(shell go env GOOS)
GOARCH := $(shell go env GOARCH)

# A helper to protect calls that push things upstreams (.e.g docker push or github artifact publish)
# $(1) - the actual command to run, if ALLOW_PUSH is not set we'll prefix this with '#' to prevent execution
define GATE_PUSH
$(if $(ALLOW_PUSH),$(1), # $(1))
endef

# The e2e tests depend on Kind kubeconfigs being in this directory,
# so this is location should not be changed by developers.
KUBECONFIG_DIR := $(HOME)/.kube

PROTOS_DEPS_PATH=$(CI_TOOLS_DIR)/protos

CLANG_FORMAT=$(CI_TOOLS_BIN_DIR)/clang-format
HELM=$(CI_TOOLS_BIN_DIR)/helm
K3D_BIN=$(CI_TOOLS_BIN_DIR)/k3d
KIND=$(CI_TOOLS_BIN_DIR)/kind
KUBEBUILDER=$(CI_TOOLS_BIN_DIR)/kubebuilder
KUBEBUILDER_ASSETS=$(CI_TOOLS_BIN_DIR)
CONTROLLER_GEN=$(CI_TOOLS_BIN_DIR)/controller-gen
KUBECTL=$(CI_TOOLS_BIN_DIR)/kubectl
PROTOC_BIN=$(CI_TOOLS_BIN_DIR)/protoc
SHELLCHECK=$(CI_TOOLS_BIN_DIR)/shellcheck
CONTAINER_STRUCTURE_TEST=$(CI_TOOLS_BIN_DIR)/container-structure-test
# from go-deps
PROTOC_GEN_GO=$(CI_TOOLS_BIN_DIR)/protoc-gen-go
PROTOC_GEN_GO_GRPC=$(CI_TOOLS_BIN_DIR)/protoc-gen-go-grpc
PROTOC_GEN_VALIDATE=$(CI_TOOLS_BIN_DIR)/protoc-gen-validate
PROTOC_GEN_JSONSCHEMA=$(CI_TOOLS_BIN_DIR)/protoc-gen-jsonschema
GINKGO=$(CI_TOOLS_BIN_DIR)/ginkgo
GOLANGCI_LINT=$(CI_TOOLS_BIN_DIR)/golangci-lint
HELM_DOCS=$(CI_TOOLS_BIN_DIR)/helm-docs
KUBE_LINTER=$(CI_TOOLS_BIN_DIR)/kube-linter
HADOLINT=$(CI_TOOLS_BIN_DIR)/hadolint
IMPORTFORMATTER=$(CI_TOOLS_BIN_DIR)/imports-formatter

TOOLS_DEPS_DIRS=$(DUBBO_DIR)/mk/dependencies
TOOLS_DEPS_LOCK_FILE=mk/dependencies/deps.lock
TOOLS_MAKEFILE=$(DUBBO_DIR)/mk/dev.mk

# Install all dependencies on tools and protobuf files
# We add one script per tool in the `mk/dependencies` folder. Add a VARIABLE for each binary and use this everywhere in Makefiles
# ideally the tool should be idempotent to make things quick to rerun.
# it's important that everything lands in $(CI_TOOLS_DIR) to be able to cache this folder in CI and speed up the build.
.PHONY: dev/tools
dev/tools: ## Bootstrap: Install all development tools
	$(TOOLS_DIR)/dev/install-dev-tools.sh $(CI_TOOLS_BIN_DIR) $(CI_TOOLS_DIR) "$(TOOLS_DEPS_DIRS)" $(TOOLS_DEPS_LOCK_FILE) $(GOOS) $(GOARCH) $(TOOLS_MAKEFILE)

.PHONY: dev/tools/clean
dev/tools/clean: ## Bootstrap: Remove all development tools
	rm -rf $(CI_TOOLS_DIR)

.PHONY: clean
clean: clean/build clean/generated clean/docs ## Dev: Clean
