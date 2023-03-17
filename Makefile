GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GO_MOD=$(GOCMD) mod
GO_ENV=$(GOCMD) env
BINARY_NAME=project-name
BINARY_UNIX=$(BINARY_NAME)-unix

all: get_vendor build

GITTAG=`git describe --tags`
GITHASH=`git rev-parse --short HEAD`

LDFLAGS=-ldflags "-X rest.GitTag=${GITTAG} -X rest.GitHash=${GITHASH}"

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
