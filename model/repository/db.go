package repository

import (
	"github.com/reaperhero/go-gin-websocket/model"
	"github.com/sirupsen/logrus"
	"strconv"
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

func (db *dbRepository) GetMessageByRoomId(roomId string, offset int) []map[string]interface{} {
	sql := "select messages.*,`users`.`username`,`users`.`avatar_id` from messages inner join users " +
		"on users.id = messages.user_id where messages.room_id = ? order by messages.id desc limit ?"
	rows, err := db.repo.Query(sql, roomId, offset)
	if err != nil {
		return nil
	}
	columns, _ := rows.Columns()
	columnLength := len(columns)
	cache := make([]interface{}, columnLength)
	for index, _ := range cache {
		var a interface{}
		cache[index] = &a
	}
	var list []map[string]interface{}
	for rows.Next() {
		_ = rows.Scan(cache...)
		item := make(map[string]interface{})
		for i, data := range cache {
			item[columns[i]] = *data.(*interface{})
		}
		list = append(list, item)
	}
	_ = rows.Close()
	logrus.Info("[dbRepository.GetMessageByRoomId]", list)
	return list
}

func (db *dbRepository) GetLimitPrivateMsg(uid, toUId string, offset int) []map[string]interface{} {
	var results []map[string]interface{}
	sql := "select messages.*, users.username,users.avatar_id from messages " +
		"inner join users on users.id = messages.user_id " +
		"where messages.user_id = ?  and messages.to_user_id= ? or messages.user_id = ? and messages.to_user_id= ? order by messages.id desc limit ?"
	err := db.repo.Select(&results, sql, uid, toUId, toUId, uid, offset)
	if err != nil {
		logrus.Info("[dbRepository.GetLimitPrivateMsg]", err, results)
	}
	return results
}

func (db *dbRepository) SaveMessageContent(content map[string]interface{}) error {
	userId := content["user_id"].(int)
	to_userId := content["to_user_id"].(int)
	text := content["content"].(string)
	room_id, _ := strconv.Atoi(content["room_id"].(string))
	imageUrl := content["image_url"].(string)
	sql := "insert into messages(`user_id`,`room_id`,`to_user_id`,`content`,`image_url`) values(?,?,?,?,?)"
	result, err := db.repo.Exec(sql, userId, room_id, to_userId, text, imageUrl)
	logrus.Info(result.LastInsertId())
	return err
}
