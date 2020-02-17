// 自动化证书配置
package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		c.String(200, "hello test")
	})
	autotls.Run(r, "www.itpp.tk")
}
