package http

import (
	"github.com/gin-gonic/gin"
	"github.com/reaperhero/go-gin-websocket/middleware"
)

type handler struct {
}

func RegisterHttpHandler(engine *gin.Engine) {
	engine.LoadHTMLGlob("views/*")
	engine.Static("/static", "./static")

	handler := handler{}
	home := engine.Group("/", middleware.EnableCookieSession())
	{
		home.GET("/", handler.index)
		home.POST("/test", handler.test)
		home.POST("/user/register", handler.userRegister)
	}
}
