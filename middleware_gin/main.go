package main

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// 日志打印到 文件里面
	if f, err := os.Create("gin.log"); err == nil {
		gin.DefaultWriter = io.MultiWriter(f)
		gin.DefaultErrorWriter = io.MultiWriter(f)
	}
	// 带中间件方式创建实例
	r := gin.New()
	// gin.Recovery() 防止服务进程挂掉
	r.Use(gin.Logger(), gin.Recovery())
	r.GET("/test", func(c *gin.Context) {
		name := c.DefaultQuery("name", "default_name")
		panic("test panic")
		c.String(200, "%s", name)
	})
	r.Run()
}
