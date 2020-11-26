package main

import (
	"github.com/gin-gonic/gin"
	"github.com/reaperhero/go-gin-websocket/handler/http"
	"github.com/reaperhero/go-gin-websocket/model/repository"
	"github.com/reaperhero/go-gin-websocket/model/usecase"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

func main() {
	repo := repository.NewRepository()
	usecase := usecase.NewUsecase(repo)
	httpRun(usecase)
}

func httpRun(usecase usecase.Usecase) {
	server := gin.Default()
	http.RegisterHttpHandler(server, usecase)
	server.Run(":8000")
}
