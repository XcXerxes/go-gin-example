package main

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/test", func(c *gin.Context) {
		// 读取所有的内容
		bodyByte, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			c.Abort()
		}
		// read之后可以通过 post 拿到值
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyByte))
		firstName := c.PostForm("first_name")
		lastName := c.DefaultPostForm("last_name", "default_last_name")
		c.String(http.StatusOK, "%s,%s,%s", firstName, lastName, string(bodyByte))
	})
	r.Run()
}
