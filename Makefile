GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GO_MOD=$(GOCMD) mod
GO_ENV=$(GOCMD) env
BINARY_NAME=iobscan-explorer-backend
BINARY_UNIX=$(BINARY_NAME)-unix
export GO111MODULE = on
export GOSUMDB=off
export GIT_TERMINAL_PROMPT=1
export GOPROXY=https://goproxy.cn,direct

all: get_vendor build

get_vendor:
	@rm -rf vendor/
	@echo "--> Downloading dependencies"
	$(GO_MOD) download
	$(GO_MOD) vendor

build:
	$(GOBUILD) -o $(BINARY_NAME) .

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)

run:
	$(GOBUILD) -o $(BINARY_NAME) -v
	./$(BINARY_NAME)


# Cross compilation
build-linux:
