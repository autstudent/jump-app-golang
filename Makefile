BINARY_PATH=bin
CMD_MAIN_PATH=cmd/golang-demo/main.go
CMD_BINARY_PATH=$(BINARY_PATH)/golang-demo
GOCMD=go
GOTEST=$(GOCMD) test
GOCLEAN=$(GOCMD) clean
GOBUILD=$(GOCMD) build

.PHONY : all test build clean run

all: test build

test:
	${GOTEST} ./...

test-verbose:
	${GOTEST} ./... -v

test-cover:
	${GOTEST} ./... -cover

build:
	$(GOBUILD) -o $(CMD_BINARY_PATH) -race $(CMD_MAIN_PATH)

clean:
	$(GOCLEAN) $(CMD_MAIN_PATH)
	rm -f $(BINARY_PATH)/*

run:
	$(CMD_BINARY_PATH)
