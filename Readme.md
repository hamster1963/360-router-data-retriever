# 360路由器后台数据获取器

该项目是一个使用Go语言编写的工具，可以用于逆向破解登录360路由器后台并获取数据。

## 功能

- 登录360路由器后台并获取数据
- 通过HTTP协议向路由器发送请求
- 解析json以获取所需的数据

## 安装

1. 克隆该项目到本地：
   `git clone https://github.com/hamster1963/360-router-data-retriever.git`

2. 进入项目目录：
  ` cd 360-router-data-retriever`

3. 安装依赖项：
   `go mod tiny`


## 使用

1. 在config/router_config.go文件中设置路由器IP地址和密码。

2. 运行main.go文件：
`go run main.go`

3. 程序将会登录路由器后台并获取数据。

## 贡献

欢迎您向该项目做出贡献！如果您发现了错误或有建议，请在GitHub上提交issue或pull request。

## 许可证

该项目采用MIT许可证。请查看LICENSE文件以获取更多信息。


