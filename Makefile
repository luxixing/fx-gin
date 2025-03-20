# Makefile
.PHONY: all build clean run test lint swagger package help

# 变量定义
APP_NAME=fix-gin
BUILD_DIR=./build
MAIN_FILE=./cmd/server/main.go
SWAGGER_FILE=./cmd/swagger/main.go

# 默认目标
all: clean lint test build

# 编译应用
build:
	@echo "编译应用..."
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(APP_NAME) $(MAIN_FILE)
	@echo "编译完成: $(BUILD_DIR)/$(APP_NAME)"

# 清理构建文件
clean:
	@echo "清理构建文件..."
	@rm -rf $(BUILD_DIR)
	@go clean
	@echo "清理完成"

# 运行应用
run:
	@echo "运行应用..."
	@go run $(MAIN_FILE)

# 运行测试
test:
	@echo "运行测试..."
	@go test -v ./...

# 代码检查
lint:
	@echo "代码检查..."
	@golangci-lint run

# 生成Swagger文档
swagger:
	@echo "检查 swag 是否已安装..."
	@if ! command -v swag > /dev/null; then \
		echo "Installing swag..."; \
		go install github.com/swaggo/swag/cmd/swag@latest; \
	fi
	@echo "生成Swagger文档..."
	@PATH="$(shell go env GOPATH)/bin:$(PATH)" go run $(SWAGGER_FILE)
# 打包项目 (忽略 .env 和 *.db 文件)
package:
	@echo "打包项目..."
	PACKAGE_NAME="fx-gin.tar.gz" && \
	echo "正在创建 $$PACKAGE_NAME..." && \
	tar --exclude='.git' --exclude='.env' --exclude='*.db' --exclude='build' \
		--exclude='*.tar.gz' --exclude='*.log' --exclude='tmp' --exclude='vendor' \
		-czf "$$PACKAGE_NAME" . && \
	echo "项目打包完成: $$PACKAGE_NAME"
# 帮助信息
help:
	@echo "可用的命令:"
	@echo "  make build          - 编译应用"
	@echo "  make clean          - 清理构建文件"
	@echo "  make run            - 运行应用"
	@echo "  make test           - 运行测试"
	@echo "  make lint           - 代码检查"
	@echo "  make swagger        - 生成Swagger文档"
	@echo "  make package        - 打包项目"
	@echo "  make all            - 执行clean, lint, test, build"
	@echo "  make help           - 显示帮助信息"