package main

import (
	"github.com/gin-gonic/gin"
	"github.com/reaperhero/go-gin-websocket/handler/http"
)

func main() {
	server := gin.Default()
	http.RegisterHttpHandler(server)
	server.Run(":8000")
}
