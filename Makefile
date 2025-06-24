# Go 参数
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod

# 项目参数
BINARY_NAME=gin-weather
BINARY_UNIX=$(BINARY_NAME)_unix
MAIN_PATH=cmd/server/main.go

# 构建目录
BUILD_DIR=bin

.PHONY: all build clean test coverage deps run docker-build docker-run help

# 默认目标
all: clean deps test build

# 构建项目
build:
	@echo "构建项目..."
	@mkdir -p $(BUILD_DIR)
	$(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME) -v $(MAIN_PATH)

# 构建 Linux 版本
build-linux:
	@echo "构建 Linux 版本..."
	@mkdir -p $(BUILD_DIR)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BUILD_DIR)/$(BINARY_UNIX) -v $(MAIN_PATH)

# 清理构建文件
clean:
	@echo "清理构建文件..."
	$(GOCLEAN)
	rm -rf $(BUILD_DIR)

# 运行测试
test:
	@echo "运行测试..."
	$(GOTEST) -v ./...

# 运行测试并生成覆盖率报告
coverage:
	@echo "生成测试覆盖率报告..."
	$(GOTEST) -coverprofile=coverage.out ./...
	$(GOCMD) tool cover -html=coverage.out -o coverage.html
	@echo "覆盖率报告已生成: coverage.html"

# 安装依赖
deps:
	@echo "安装依赖..."
	$(GOMOD) download
	$(GOMOD) tidy

# 运行项目
run:
	@echo "运行项目..."
	$(GOCMD) run $(MAIN_PATH)

# 格式化代码
fmt:
	@echo "格式化代码..."
	$(GOCMD) fmt ./...

# 代码检查
vet:
	@echo "代码检查..."
	$(GOCMD) vet ./...

# 构建 Docker 镜像
docker-build:
	@echo "构建 Docker 镜像..."
	docker build -t $(BINARY_NAME):latest .

# 运行 Docker 容器
docker-run:
	@echo "运行 Docker 容器..."
	docker run -p 8080:8080 --env-file .env $(BINARY_NAME):latest

# 使用 docker-compose 启动
docker-compose-up:
	@echo "使用 docker-compose 启动服务..."
	docker-compose up -d

# 使用 docker-compose 停止
docker-compose-down:
	@echo "使用 docker-compose 停止服务..."
	docker-compose down

# 查看 docker-compose 日志
docker-compose-logs:
	@echo "查看 docker-compose 日志..."
	docker-compose logs -f

# 安装开发工具
install-tools:
	@echo "安装开发工具..."
	$(GOGET) -u golang.org/x/tools/cmd/goimports
	$(GOGET) -u github.com/golangci/golangci-lint/cmd/golangci-lint

# 代码质量检查
lint:
	@echo "运行代码质量检查..."
	golangci-lint run

# 生成 API 文档（如果使用 swag）
# docs:
# 	@echo "生成 API 文档..."
# 	swag init -g cmd/server/main.go

# 显示帮助信息
help:
	@echo "可用的 make 命令:"
	@echo "  build          - 构建项目"
	@echo "  build-linux    - 构建 Linux 版本"
	@echo "  clean          - 清理构建文件"
	@echo "  test           - 运行测试"
	@echo "  coverage       - 生成测试覆盖率报告"
	@echo "  deps           - 安装依赖"
	@echo "  run            - 运行项目"
	@echo "  fmt            - 格式化代码"
	@echo "  vet            - 代码检查"
	@echo "  docker-build   - 构建 Docker 镜像"
	@echo "  docker-run     - 运行 Docker 容器"
	@echo "  docker-compose-up   - 使用 docker-compose 启动"
	@echo "  docker-compose-down - 使用 docker-compose 停止"
	@echo "  docker-compose-logs - 查看 docker-compose 日志"
	@echo "  install-tools  - 安装开发工具"
	@echo "  lint           - 代码质量检查"
	@echo "  help           - 显示此帮助信息"
