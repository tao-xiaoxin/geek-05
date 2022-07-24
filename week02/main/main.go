package main

import (
	"github.com/gin-gonic/gin"
	"wqhn.cn/geekbang-homework-module2/controller"
)

// 编写一个 HTTP 服务器：
// 1 接收客户端 request，并将 request 中带的 header 写入 response header
// 2 读取当前系统的环境变量中的 VERSION 配置，并写入 response header
// 3 Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
// 4 当访问 localhost/healthz 时，应返回 200
func main(){

	gin.SetMode("release")  // 设置为生产模式

	router := gin.New()
	router.GET("/header", controller.Header)
	router.GET("/version", controller.Version)
	router.GET("/clientip", controller.ClientIp)
	router.GET("/healthz", controller.Healthz)

	router.Run(":80")
}

