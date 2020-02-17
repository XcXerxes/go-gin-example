package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {

	})
	srv := &http.Server{Addr: ":8085", Handler: r}

	go func ()  {
		if err := srv.ListenAndServe(); err != nil && err !=http.ErrServerClosed{log.Fatal(err)}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGALRM)
	<-quit
	log.Println("shutdown")
	
}