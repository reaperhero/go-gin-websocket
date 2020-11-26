package usecase

import (
	"errors"
	"github.com/reaperhero/go-gin-websocket/model"
	"github.com/reaperhero/go-gin-websocket/utils"
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
