package main

import (
	"github.com/gin-gonic/gin"
	handlersfunc "go-file-upload/handlers"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("index.html")
	r.GET("/", handlersfunc.HandleHome)
	r.POST("/upload", handlersfunc.HandleUpload)

	r.Run(":9020")
}