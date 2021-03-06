package usecase

import (
	"github.com/reaperhero/go-gin-websocket/model"
	"github.com/reaperhero/go-gin-websocket/model/repository"
)

type Usecase interface {
	SaveUser(user model.User) error
	FindUserByName(username string) model.User
	GetMessageByRoomId(roomId string) []map[string]interface{}
	SaveMessageContent(message map[string]interface{})
	GetLimitPrivateMsg(uid, toUId string) []map[string]interface{}
}

type useacse struct {
	repo repository.Repository
}

func NewUsecase(repo repository.Repository) Usecase {
	return &useacse{repo: repo}
}
