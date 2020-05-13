# 淘宝 TOP SDK
top sdk for golang

## Install

1. 配置 gitlab

```shell
git config --global url.git@gitlab.51h5.com:2202.insteadOf https://gitlab.51h5.com
```

2. 配置 golang

```shell
go env -w GO111MODULE=on
go env -w GOPRIVATE=gitlab.51h5.com
```

3. 配置 go.mod

```shell
go mod tidy
```

## Example

> see `_example/main.go`
