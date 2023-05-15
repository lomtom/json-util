# json-util

json获取工具
```shell
json-util -f 1.json -k name
```

# 安装

1. 支持Arm64和Amd64架构下载
- [Amd64](https://github.com/lomtom/json-util/releases/download/v1.0/json-util-amd64-tar.gz)
- [Arm64](https://github.com/lomtom/json-util/releases/download/v1.0/json-util-arm64-tar.gz)

2. 源码编译(需要git和go1.18以上)
```shell
git clone https://github.com/lomtom/json-util.git
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/json-util cmd/main.go
```