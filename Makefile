export GO111MODULE=on

# Binary file names.
BINARY_NAME=web-app

# Build parameters.
BUILD_PATH=".build"

.PHONY: build
build: fmt-go clean build-go

build-go:
	mkdir -p ${BUILD_PATH}
	cp -r Makefile build/*.sh cmd pkg internal configs ${BUILD_PATH}
	go build -a -o $(BUILD_PATH)/${BINARY_NAME} $(BUILD_PATH)/cmd/main.go

fmt-go:
	@go fmt ./pkg/... ./internal/... ./cmd/...
	@gofmt -w -s pkg internal cmd
	@goimports -w pkg internal cmd

tidy:
	go mod tidy

clean:
	mkdir -p ${BUILD_PATH}
	rm -rf $(BUILD_PATH)

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
