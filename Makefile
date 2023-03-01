export GO111MODULE=on

# Binary file names.
BINARY_NAME=web-app

# Build parameters.
BUILD_PATH=.build
DIST_PATH=.dist

.PHONY: build
build: build-go clean

build-go:
	mkdir -p ${BUILD_PATH}
	rm -rf ${DIST_PATH}
	cp -r Makefile build/*.sh cmd pkg internal configs ${BUILD_PATH}
	CGO_ENABLED=0 go build -mod=readonly -v -o $(DIST_PATH)/${BINARY_NAME} $(BUILD_PATH)/cmd/main.go
	cp -rf ${BUILD_PATH}/configs $(DIST_PATH)

fmt-go:
	@echo 'fmt-go'
	@go fmt ./pkg/... ./internal/... ./cmd/...
	@gofmt -w -s pkg internal cmd
	@goimports -w pkg internal cmd

tidy:
	go mod tidy

clean:
	@echo 'clean'
	@rm -rf $(BUILD_PATH)

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
