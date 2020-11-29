package http

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"github.com/reaperhero/go-gin-websocket/middleware"
	"github.com/reaperhero/go-gin-websocket/model/usecase"
)

const (
	//64位
	cookieStoreAuthKey = "4238uihfieh49"
	//AES encrypt key必须是16或者32位
	cookieStoreEncryptKey = "..."
)

type handler struct {
	usecase usecase.Usecase
}

func RegisterHttpHandler(engine *gin.Engine, usecase usecase.Usecase) {
	engine.LoadHTMLGlob("views/*")
	engine.Static("/static", "./static")

	handler := handler{
		usecase: usecase,
	}

	store, _ := redis.NewStore(10, "tcp", "localhost:6379", "Jsx4ujds2P8veOCgz", []byte(cookieStoreAuthKey))
	engine.Use(sessions.Sessions("sessionStore", store))

	home := engine.Group("/")
	{
		home.GET("/", handler.index)
		home.POST("/register", handler.userRegister)
		home.POST("/login", handler.userLogin)
	}

	authorized := engine.Group("/", middleware.AuthSessionMiddle())
	{
		authorized.GET("/logout", handler.userLogout)
		authorized.GET("/home", handler.home)
		authorized.GET("/room/:room_id", handler.room)
		authorized.GET("/private-chat", handler.privateChat)
		authorized.GET("/ws", handler.wsHandler)
		authorized.GET("/pagination", handler.pagination)
	}
}
