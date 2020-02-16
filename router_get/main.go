package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/user", func(ctx *gin.Context) {
		firstName := ctx.Query("first_name")
		lastName := ctx.DefaultQuery("last_name", "default_last_name")
		ctx.JSON(200, gin.H{"firstName": firstName, "lastName": lastName})
	})
	r.Run()
}
