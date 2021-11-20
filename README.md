# 读取 Github 关注数

## 编译

```shell
CGO_ENABLED=0 go build -ldflags="-s -w" -o fetch-github-stars main.go
```

## 执行

```shell
./fetch-github-stars hyperf/hyperf
```