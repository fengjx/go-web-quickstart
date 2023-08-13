export GO111MODULE=on

# Binary file names.
BINARY_NAME=app

# Build parameters.
BUILD_PATH=.build
DIST_PATH=.dist

.PHONY: build
### build:				项目打包
build: build-go clean

### build-go:				构建 golang 包
build-go:
	mkdir -p ${BUILD_PATH}
	rm -rf ${DIST_PATH}
	cp -r Makefile build/*.sh cmd pkg internal configs ${BUILD_PATH}
	CGO_ENABLED=0 go build -tags=jsoniter -mod=readonly -v -o $(DIST_PATH)/${BINARY_NAME} $(BUILD_PATH)/cmd/main.go
	cp -rf ${BUILD_PATH}/configs $(DIST_PATH)
	rm -rf ${BUILD_PATH}

### fmt-go:				格式化 golang 代码
fmt-go:
	@echo 'fmt-go'
	@go fmt ./pkg/... ./internal/... ./cmd/...
	@gofmt -w -s pkg internal cmd
	@goimports -w pkg internal cmd

### tidy:				去掉未使用的项目依赖
tidy:
	go mod tidy

### clean:				清理临时文件
clean:
	@echo 'clean'
	@rm -rf $(BUILD_PATH)


### help:				Makefile 帮助
.PHONY: help
help:
	@echo Makefile cmd:
	@echo
	@grep -E '^### [-A-Za-z0-9_]+:' Makefile | sed 's/###/   /'