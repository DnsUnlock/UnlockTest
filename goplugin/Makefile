# UnlockTest Go 插件 Makefile
# 提供便捷的编译和测试命令

.PHONY: build test clean help releases

# 默认目标
all: build

# 编译插件
build:
	@echo "编译 UnlockTest 插件..."
	go build -o unlock-test.tool.plugin .
	@echo "插件编译完成: unlock-test.tool.plugin"

# 运行测试
test: build
	@echo "运行插件测试..."
	go run ../test_plugin_client.go

# 清理生成的文件
clean:
	@echo "清理生成的文件..."
	rm -f unlock-test.tool.plugin
	rm -rf releases/
	@echo "清理完成"

# 安装依赖
deps:
	@echo "安装依赖包..."
	go mod tidy
	go get github.com/gophertool/tool/plugin
	@echo "依赖安装完成"

# 创建多平台发布版本
releases:
	@echo "创建多平台发布版本..."
	@mkdir -p releases
	# Linux 平台
	@mkdir -p releases/linux/amd64 releases/linux/arm64
	@echo "编译 Linux amd64..."
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o releases/linux/amd64/UnlockTest.tool.plugin .
	@echo "编译 Linux arm64..."
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o releases/linux/arm64/UnlockTest.tool.plugin .
	# Windows 平台
	@mkdir -p releases/windows/amd64 releases/windows/arm64
	@echo "编译 Windows amd64..."
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o releases/windows/amd64/UnlockTest.tool.plugin.exe .
	@echo "编译 Windows arm64..."
	CGO_ENABLED=0 GOOS=windows GOARCH=arm64 go build -o releases/windows/arm64/UnlockTest.tool.plugin.exe .
	# macOS 平台
	@mkdir -p releases/darwin/amd64 releases/darwin/arm64
	@echo "编译 macOS amd64..."
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o releases/darwin/amd64/UnlockTest.tool.plugin .
	@echo "编译 macOS arm64..."
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o releases/darwin/arm64/UnlockTest.tool.plugin .
	@echo "所有平台编译完成，文件保存在 releases/ 目录"
	@echo "文件结构: 平台/架构/UnlockTest.tool.plugin(.exe)"
	@echo "支持的平台和架构:"
	@echo "  Linux: amd64, arm64"
	@echo "  Windows: amd64, arm64 (带.exe后缀)"
	@echo "  macOS: amd64, arm64"
	@echo "目录结构:"
	@find releases -name "UnlockTest.tool.plugin*"
	@echo "总计文件数: $$(find releases -name 'UnlockTest.tool.plugin*' | wc -l)"

# 显示帮助信息
help:
	@echo "可用的命令:"
	@echo "  build    - 编译插件"
	@echo "  test     - 编译并测试插件"
	@echo "  clean    - 清理生成的文件"
	@echo "  deps     - 安装依赖包"
	@echo "  releases - 创建多平台发布版本"
	@echo "  help     - 显示此帮助信息"

# 开发模式：监听文件变化并自动重新编译
dev:
	@echo "开发模式：监听文件变化..."
	@echo "请安装 fswatch: brew install fswatch"
	fswatch -o . | xargs -n1 -I{} make build