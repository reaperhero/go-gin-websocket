package repository

import (
	"github.com/reaperhero/go-gin-websocket/model"
	"github.com/sirupsen/logrus"
	"time"
)

func (db *dbRepository) SaveUser(username, password, avatar_id string) error {
	sql := "insert into users(`username`,`password`,`avatar_id`,`created_at`,`updated_at`) values(?, ?, ?,?, ?)"
	result, err := db.repo.Exec(sql, username, password, avatar_id, time.Now(), time.Now())
	if err != nil {
		logrus.Println("exec failed, ", err)
		return err
	}
	logrus.Info(result.LastInsertId())
	return nil
}

func (db *dbRepository) FindUserById(username string) model.User {
	sql := "select username,password,avatar_id from users where username = ?"
	user := model.User{}
	err := db.repo.Get(&user, sql, username)
	if err != nil {
		logrus.Println("exec failed, ", err)
	}
	return user
}
