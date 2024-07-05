#/bin/bash

# build plugin
echo "Building plugin"
# 编译并去除符号表和调试信息
go build -ldflags "-s -w" -gcflags "all=-l -B" -o ./build/4GTV.so -buildmode=plugin ./plugin/4GTV/main.go

chmod +x ./build/4GTV.so

# 调用 strip
strip ./build/4GTV.so

# 使用 upx --brute 压缩
sudo upx --brute ./build/4GTV.so

# 使用 zstd 进行最终压缩
zstd -19 ./build/4GTV.so -o ./build/4GTV.so
