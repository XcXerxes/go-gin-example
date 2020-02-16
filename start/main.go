package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// 创建一个 gin 默认的服务
	r := gin.Default()
	// 监听路由
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong"})
	})
	// 监听默认8080端口
	r.Run() // 8080
}
