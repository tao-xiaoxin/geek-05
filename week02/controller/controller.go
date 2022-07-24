package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

// 1 接收客户端 request，并将 request 中带的 header 写入 response header
func Header(context *gin.Context){
	// context.Request.Header是一个http.Header，它是对map[string][]string的自定义类型。
	// 详见http包下的header.go文件
	for key, val := range context.Request.Header {
		context.Writer.Header().Add(key, val[0])
	}

	// 另一种写法是直接调用Header.Get方法
	//for key := range context.Request.Header {
	//	val := context.Request.Header.Get(key)
	//	context.Writer.Header().Add(key, val)
	//}
	context.String(200, "%s", "ok")
}

// 2 读取当前系统的环境变量中的 VERSION 配置，并写入 response header
func Version(context *gin.Context){
	val := os.Getenv("VERSION")
	context.Writer.Header().Set("version", val)
	context.String(200, "%s", "ok")
}

// 3 ClientIp 3.Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
func ClientIp(context *gin.Context){
	httpCode := http.StatusOK
	fmt.Printf("Client IP:%s, HTTP response Code:%d\n", context.ClientIP(), httpCode)
	context.String(httpCode, "Your IP:%s", context.ClientIP())
}

// 4 Healthz 4.当访问 localhost/healthz 时，应返回 200
func Healthz(context *gin.Context){
	context.String(200, "%s", "ok")
}
