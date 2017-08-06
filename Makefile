DBG_MAKEFILE ?=
ifeq ($(DBG_MAKEFILE),1)
    $(warning ***** starting Makefile for goal(s) "$(MAKECMDGOALS)")
    $(warning ***** $(shell date))
else
    # If we're not debugging the Makefile, don't echo recipes.
    MAKEFLAGS += -s
endif

# It's necessary to set this because some environments don't link sh -> bash.
SHELL := /bin/bash

# go option
GO        ?= go
BINDIR    := $(CURDIR)/bin

# Required for globs to work correctly
SHELL=/bin/bash

.PHONY: all
all: build ## build binary

.PHONY: build
build: verify ## build binary
	GOBIN=$(BINDIR) $(GO) install

HAS_GLIDE := $(shell command -v glide;)
HAS_GIT := $(shell command -v git;)
HAS_HG := $(shell command -v hg;)

.PHONY: bootstrap
bootstrap: ## build the deb packages
ifndef HAS_GLIDE
	go get -u github.com/Masterminds/glide
endif

ifndef HAS_GIT
	$(error You must install Git)
endif
	glide install --strip-vendor

.PHONY: clean
clean: ## clean up cached resources
	@rm -rf $(BINDIR)

.PHONY: verify
verify: ## verify source files
	hack/verify-gofmt.sh

help: ## this help
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {sub("\\\\n",sprintf("\n%22c"," "), $$2);printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)
