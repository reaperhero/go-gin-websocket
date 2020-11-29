package repository

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/reaperhero/go-gin-websocket/model"
	"github.com/reaperhero/go-gin-websocket/utils"
	"github.com/sirupsen/logrus"
	"time"
)

var (
	db *sqlx.DB
)

type Repository interface {
	SaveUser(username, password, avatarId string) error
	FindUserById(username string) model.User
	GetMessageByRoomId(roomId string, offset int) []map[string]interface{}
	SaveMessageContent(content map[string]interface{}) error
	GetLimitPrivateMsg(uid, toUId string, offset int) []map[string]interface{}
}

type dbRepository struct {
	repo *sqlx.DB
}

func init() {
	databaseHost := utils.GetEnvWithDefault("DBHOST", "127.0.0.1")
	databaseName := utils.GetEnvWithDefault("DBNAME", "ws")
	databaseUser := utils.GetEnvWithDefault("DBUSER", "root")
	databasePort := utils.GetEnvWithDefault("DBPORT", "3306")
	databasePass := utils.GetEnvWithDefault("DBPASS", "123456")

	var err error
	connStr := databaseUser + ":" + databasePass + "@tcp(" + databaseHost + ":" + databasePort + ")/" + databaseName + "?parseTime=true&charset=utf8mb4&collation=utf8mb4_unicode_ci"
	db, _ = sqlx.Open("mysql", connStr)
	logrus.Info(connStr, err)
	db.SetConnMaxLifetime(time.Minute * 10)
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	if err == nil {
		err = db.Ping()
	}

	if err != nil {
		logrus.Fatalf("database connect error: %s", err)
	}
}

func NewRepository() Repository {
	return &dbRepository{
		repo: db,
	}
}
