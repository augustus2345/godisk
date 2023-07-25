# go_disk
> 基于 go-zero 搭建的一个轻量级的云盘项目

使用到的命令
```text
// 创建 api 命令
goctl api new core // 使用 goctl 脚本服务新建一个 core

// 使用 api 生成代码命令
goctl api go -api core.api -dir . -style go_zero

// 启动服务
go run core.go -f etc/core-api.yaml
```