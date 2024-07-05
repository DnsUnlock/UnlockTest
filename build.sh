#/bin/bash

echo "拉取 Docker 镜像"
docker pull dnsunlockcom/build-go-1.22.5:cn
echo "构建 Docker 镜像"
docker build -t my-golang-app .
echo "运行 Docker 容器"
docker run --name temp-container my-golang-app
echo "复制生成的动态库到当前目录"
docker cp temp-container:/app/unlock_test.so ./unlock_test.so
echo "删除 Docker 容器"
docker rm temp-container
echo "删除 Docker 镜像"
docker rmi my-golang-app
