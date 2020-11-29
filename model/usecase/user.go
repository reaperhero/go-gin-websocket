package usecase

import (
	"errors"
	"github.com/reaperhero/go-gin-websocket/model"
	"github.com/reaperhero/go-gin-websocket/utils"
	"github.com/sirupsen/logrus"
)

func (u *useacse) SaveUser(user model.User) error {
	hashpassword := utils.GenerateFromPassword(user.Password)
	err := u.repo.SaveUser(user.Username, hashpassword, user.AvatarId)
	if err != nil {
		return errors.New("用户保存失败")
	}
	return nil
}

func (u *useacse) FindUserByName(username string) model.User {
	return u.repo.FindUserById(username)
}

func (u *useacse) GetMessageByRoomId(roomId string) []interface{} {

	list := u.repo.GetMessageByRoomId(roomId, 100)
	if list == nil {
		logrus.Println("[useacse.GetMessageByRoomId]", nil)
		return nil
	}
	return list
}
