package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// 1
	r.Static("/assets", "./assets")
	// 2
	r.StaticFS("/static", http.Dir("static"))
	// 3
	r.StaticFile("/favicon.ico", "./favicon.ico")
	r.Run()
}
