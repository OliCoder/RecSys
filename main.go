package main

import (
	"flag"
	"github.com/gin-gonic/gin"
)

func main() {
	flag.Parse()
	router := gin.Default()
	router.GET("/")
	router.POST("/")
	router.Run()
}
