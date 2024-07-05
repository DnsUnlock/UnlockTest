#/bin/bash

# build plugin
echo "Building plugin"
# 编译并去除符号表和调试信息
go build -ldflags "-s -w" -o ./build/unlock_test.so -buildmode=plugin ./plugin/main.go