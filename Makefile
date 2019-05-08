# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GORELEASERCMD=goreleaser

# Binary names
BINARY_NAME=terraform-provider-genymotion

all: clean build testacc release

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

clean: ## Cleans up generated files/folders from the build.
	rm -rfv "dist/" "${BINARY_NAME}"

build: clean ## Builds the application.
	$(GOBUILD) -o $(BINARY_NAME) -v

testacc: clean build ## Run acceptance tests
	TF_ACC=1 go test -v ./genymotion -timeout 120m

release: clean build testacc ## Generate a release, but don't publish to GitHub.
	$(GORELEASERCMD) --skip-validate --skip-publish
