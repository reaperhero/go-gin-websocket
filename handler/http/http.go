package http

import "github.com/gin-gonic/gin"

type handler struct {
}

func RegisterHttpHandler(engine *gin.Engine) {

	home := engine.Group("/")
	{
		home.GET("/", func(context *gin.Context) {
			context.String(200, "ok")
			return
		})

	}
}
