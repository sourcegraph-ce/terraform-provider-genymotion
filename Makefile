# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

# Binary names
BINARY_NAME=terraform-provider-genymotion

all: clean build testacc

clean:
	rm -f $(BINARY_NAME)

build:
	$(GOBUILD) -o $(BINARY_NAME) -v

testacc: 
	TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout 120m
