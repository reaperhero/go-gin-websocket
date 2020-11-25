package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	httphandler "github.com/reaperhero/go-gin-websocket/handler/http"
	"net/http"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func ping(c *gin.Context) {
	//升级get请求为webSocket协议
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer ws.Close()
	for {
		//读取ws中的数据
		mt, message, err := ws.ReadMessage()
		if err != nil {
			break
		}
		if string(message) == "ping" {
			message = []byte("pong")
		}
		//写入ws数据
		err = ws.WriteMessage(mt, message)
		if err != nil {
			break
		}
	}
}

func main() {
	server := gin.Default()
	httphandler.RegisterHttpHandler(server)
	server.Run(":8000")
}
