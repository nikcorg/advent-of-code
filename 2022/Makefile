SHELL = /bin/bash

TEST_TARGETS = $(shell find . -type d -maxdepth 1 -mindepth 1 ! -name "bin" | sort)
BUILD_TARGETS = $(shell find . -type d -maxdepth 1 -mindepth 1 ! -name "bin" ! -name "util" | sort)

target = $(word 1, $@)

.PHONY: test
test:
	go test $(foreach trg,$(TEST_TARGETS),$(trg)/...)

.PHONY: test
cover:
	go test -coverprofile=coverage.out $(foreach trg,$(TEST_TARGETS),$(trg)/...) && \
	go tool cover -html=coverage.out

.PHONY: $(BUILD_TARGETS)
$(BUILD_TARGETS):
	go test ./$(target)/... && go build -o ./bin/solve$(target) ./$(target)

.PHONY: all
all: $(BUILD_TARGETS)

.PHONY: clean
clean:
	rm -fr ./bin coverage.out
