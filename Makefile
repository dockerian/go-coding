# Go Makefile for go-coding
# Set OS platform
# See http://stackoverflow.com/questions/714100/os-detecting-makefile
# TODO: macro commands 'cp', 'mkdir', 'mv', 'rm', etc. for Windows
ifeq ($(shell uname),Darwin) # Mac OS
	OS_PLATFORM := darwin
	OS_PLATFORM_NAME := Mac OS
else ifeq ($(OS),Windows_NT) # Windows
		OS_PLATFORM := windows
		OS_PLATFORM_NAME := Windows
else
	OS_PLATFORM := linux
	OS_PLATFORM_NAME := Linux
endif

# Set build parameters
BINARY ?= go-coding
BUILD_OS ?= $(OS_PLATFORM)
BUILD_VERSION ?= $(shell cat release/tag)
BUILD_MASTER_VERSION ?= 0
BUILD_PREFIX := $(BINARY)-$(BUILD_VERSION)
OWNER := dockerian
PROJECT := go-coding
PROJECT_PACKAGE := github.com/$(OWNER)/$(PROJECT)
CMD_PACKAGE := $(PROJECT_PACKAGE)/cli/cmd
SOURCE_PATH := $(GOPATH)/src/github.com/$(OWNER)/$(PROJECT)

HUB_ACCOUNT := dockerian
DOCKER_IMAG := go-coding
DOCKER_TAGS := $(HUB_ACCOUNT)/$(DOCKER_IMAG)

# TODO: Test the Makefile macros
# to represent "ifdef VAR1 || VAR2", use
#		ifneq ($(call ifdef_any,VAR1 VAR2),) # ifneq ($(VAR1)$(VAR2),)
# to represent "ifdef VAR1 && VAR2", use
#		ifeq ($(call ifdef_none,VAR1 VAR2),) # ifneq ($(and $(VAR1),$(VAR2)),)
ifdef_any := $(filter-out undefined,$(foreach v,$(1),$(origin $(v))))
ifdef_none := $(filter undefined,$(foreach v,$(1),$(origin $(v))))

# Set testing parameters
TEST_MATCH ?= .
ifndef TEST_TAGS
	TEST_TAGS := all
endif
ifdef TEST_DIR
	TEST_COVERAGE := -cover -coverprofile cover.out ./$(TEST_DIR)
	TEST_PACKAGE := $(PROJECT_PACKAGE)/$(TEST_DIR)
else
	TEST_PACKAGE := ./... # $(shell go list ./... | grep -v /vendor/)
endif
ifneq ($(TEST_VERBOSE)$(VERBOSE),)
	TEST_VERBOSE := -v
endif
TEST_ARGS := $(TEST_VERBOSE) -bench='$(TEST_MATCH)' -run='$(TEST_MATCH)' -tags='$(TEST_TAGS)' $(TEST_COVERAGE)

# Set the -ldflags option for go build, interpolate the variable values
LDFLAGS := -ldflags "-X '$(PROJECT_PACKAGE).buildVersion=$(BUILD_VERSION)'"

# Set variables for distribution
DIST_ARCH := amd64
DIST_DIR := dist
DIST_DOWNLOADS := $(DIST_DIR)/downloads
DIST_UPDATES := $(DIST_DIR)/v$(BUILD_MASTER_VERSION)
DIST_VER := $(DIST_UPDATES)/$(BUILD_VERSION)
DIST_PREFIX := $(DIST_DOWNLOADS)/$(BUILD_PREFIX)
GO_SELF_UPDATE_INPUTS := $(SOURCE_PATH)/build/updates
GO_SELF_UPDATE_PUBLIC := $(SOURCE_PATH)/public
BUILD_DIR := build
BIN_DIR := $(BUILD_DIR)/bin


.PHONY: build build-all clean cmd qb run test fmt lint list vet

default: cmd
all: build-all run test

build: clean fmt lint vet show-env
	@echo "............................................................"
	@echo "Building: '$(BINARY)' ... [BUILD_OS = $(BUILD_OS)]"
	go get -u
	go get -u github.com/tools/godep
	godep restore

	GOARCH=$(DIST_ARCH) GOOS=$(BUILD_OS) go build $(LDFLAGS) -o $(BIN_DIR)/$(BUILD_OS)/$(BINARY) main.go

	@echo ""
	@echo "Copying $(BIN_DIR)/$(BUILD_OS)/$(BINARY) [BUILD_OS = $(BUILD_OS)]"
	cp -f $(BIN_DIR)/$(BUILD_OS)/$(BINARY) $(BUILD_DIR)/$(BINARY)
	@echo "DONE: [$@]"

build-all: clean fmt lint vet show-env qb
	@echo "............................................................"
	@echo "Building $(BINARY) for all platforms..."
	go get -u
	go get -u github.com/tools/godep
	godep restore
	go get -t github.com/sanbornm/go-selfupdate
	go install github.com/sanbornm/go-selfupdate

	mkdir -p $(BIN_DIR)
	mkdir -p $(DIST_DOWNLOADS)
	mkdir -p $(GO_SELF_UPDATE_INPUTS)

	@- $(foreach os,darwin linux windows, \
		echo ""; \
		echo "Building $(BUILD_VERSION) for $(os) platform"; \
		echo "GOARCH=$(DIST_ARCH) GOOS=$(os) go build $(LDFLAGS) -o $(BIN_DIR)/$(os)/$(BINARY) main.go"; \
		GOARCH=$(DIST_ARCH) GOOS=$(os) go build $(LDFLAGS) -o $(BIN_DIR)/$(os)/$(BINARY) main.go; \
		cp -p $(BIN_DIR)/$(os)/$(BINARY) $(GO_SELF_UPDATE_INPUTS)/$(os)-$(DIST_ARCH); \
		if [[ "$(os)" == "windows" ]]; then \
			mv $(BIN_DIR)/$(os)/$(BINARY) $(BIN_DIR)/$(os)/$(BINARY).exe; \
			zip -jr $(DIST_PREFIX)-$(os)-$(DIST_ARCH).zip $(BIN_DIR)/$(os)/$(BINARY).exe; \
		else \
			tar -C $(BIN_DIR)/$(os)/ -cvzf $(DIST_PREFIX)-$(os)-$(DIST_ARCH).tar.gz ./$(BINARY); \
		fi; \
	)
	@echo ""

	# create self-update distribution in public folder
	go-selfupdate "$(GO_SELF_UPDATE_INPUTS)" "$(BUILD_VERSION)"

	mkdir -p "$(DIST_VER)"
	rm -rf "$(DIST_VER)"
	cp -rf "$(GO_SELF_UPDATE_PUBLIC)"/* "$(DIST_UPDATES)/"
	cp -rf "$(GO_SELF_UPDATE_PUBLIC)"/*.json "$(DIST_VER)/"
	rm -rf "$(GO_SELF_UPDATE_PUBLIC)"

	# show distribution
	@tree "$(DIST_DIR)" 2>/dev/null; true
	@echo "DONE: $@"

qb:
	@echo "............................................................"
	@echo "Building $(BIN_DIR)/$(OS_PLATFORM)/$(BINARY) [OS = $(OS_PLATFORM)]"
	GOARCH=$(DIST_ARCH) GOOS=$(OS_PLATFORM) go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY) main.go
	@echo "DONE: [$@]"

cmd:
	@echo "............................................................"
	@echo "Start cmd shell in docker container ..."
	./run.sh "cmd"
	@echo "DONE: [$@]"

run:
	@echo "............................................................"
	@echo "Running: $(BUILD_DIR)/$(BINARY) ..."
	@$(BUILD_DIR)/$(BINARY)
	@echo "DONE: [$@]"

test:
	@echo "............................................................"
	@echo "Running tests ... [tags: $(TEST_TAGS)]"
	@echo "go test $(TEST_PACKAGE) $(TEST_ARGS)"
	go test $(TEST_PACKAGE) $(TEST_ARGS) 2>&1 | tee ./tests.log
	@tools/get-test-summary.sh "./tests.log"
	@echo "DONE: [$@]"


clean:
	@echo "============================================================"
	@echo "Cleaning build..."
	rm -rf $(BIN_DIR)
	rm -rf $(BUILD_DIR)
	rm -rf $(DIST_DIR)
	@echo "DONE: [$@]"


clean-all: clean
ifeq ("$(wildcard /.dockerenv)","")
	# not in a docker container
	docker rm -f $(docker ps -a|grep ${DOCKER_IMAG}|awk '{print $1}') 2>/dev/null || true
	docker rmi -f $(docker images -a|grep ${DOCKER_TAGS} 2>&1|awk '{print $1}') 2>/dev/null || true
endif
	@echo "DONE: [$@]"


show-env:
	@echo "............................................................"
	@echo "OS Platform: "$(OS_PLATFORM_NAME)
	@echo "------------------------------------------------------------"
	@echo "GOPATH = $(GOPATH)"
	@echo "GOROOT = $(GOROOT)"
	@echo " SHELL = $(SHELL)"
	@echo ""
	@env | sort
	@echo ""


# http://golang.org/cmd/go/#hdr-Run_gofmt_on_package_sources
fmt:
	@echo "Formatting code..."
	go fmt ./...
	@echo "DONE: [$@]"

lint:
	@echo "Check coding style..."
	go get -u github.com/golang/lint/golint
	golint -set_exit_status puzzle
	@echo "DONE: [$@]"

list:
	@$(MAKE) -pRrq -f $(lastword $(MAKEFILE_LIST)) : 2>/dev/null | awk -v RS= -F: '/^# File/,/^# Finished Make data base/ {if ($$1 !~ "^[#.]") {print $$1}}' | sort | egrep -v -e '^[^[:alnum:]]' -e '^$@$$' | xargs

# http://godoc.org/code.google.com/p/go.tools/cmd/vet
# go get code.google.com/p/go.tools/cmd/vet
vet:
	@echo "Check go code correctness..."
	go vet ./...
	@echo "DONE: [$@]"
