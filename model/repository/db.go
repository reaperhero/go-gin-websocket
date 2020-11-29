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
	sql := "select id,username,password,avatar_id from users where username = ?"
	user := model.User{}
	err := db.repo.Get(&user, sql, username)
	if err != nil {
		logrus.Println("exec failed, ", err)
	}
	return user
}

func (db *dbRepository) GetMessageByRoomId(roomId string, offset int) []interface{} {
	sql := "select messages.*,users.username ,users.avatar_id from messages inner join users on users.id = messages.user_id where messages.room_id = ? order by messages.id desc limit ?"
	result := make([]*struct {
		ID        uint       `db:"id" json:"id"`
		UserId    int        `db:"user_id" json:"user_id"`
		ToUserId  int        `db:"to_user_id" json:"to_user_id"`
		RoomId    int        `db:"room_id" json:"room_id"`
		Content   string     `db:"content" json:"content"`
		ImageUrl  string     `db:"image_url" json:"image_url"`
		Username  string     `db:"username" json:"username"`
		AvatarId  int        `db:"avatar_id" json:"avatar_id"`
		CreatedAt *time.Time `db:"created_at" json:"created_at"`
		UpdatedAt *time.Time `db:"updated_at" json:"updated_at"`
		DeletedAt *time.Time `db:"deleted_at" json:"deleted_at"`
	}, 0)
	err := db.repo.Select(&result, sql, roomId, offset)
	if err != nil {
		logrus.Println("[dbRepository.GetMessageByRoomId]", err)
		return nil
	}
	var interfaceSlice = make([]interface{}, len(result))
	for i, d := range result {
		interfaceSlice[i] = d
	}
	return interfaceSlice
}
