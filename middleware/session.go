package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/reaperhero/go-gin-websocket/model"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func SaveAuthSession(c *gin.Context, info interface{}) { // info不能为struct
	session := sessions.Default(c)
	session.Set("uid", info)
	err := session.Save()
	if err != nil {
		logrus.WithField("[SaveAuthSession]:", err).Info(err)
	}
}

func GetSessionUserInfo(c *gin.Context) *model.User {
	session := sessions.Default(c)

	uid := session.Get("uid")
	user, ok := uid.(model.User)
	logrus.WithField("user", user).Info("[session GetSessionUserInfo]")
	if !ok {
		logrus.WithField("session", user).Info("session nil")
		return nil
	}
	return &user
}

func ClearAuthSession(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
}

func HasSession(c *gin.Context) bool {
	session := sessions.Default(c)
	if sessionValue := session.Get("uid"); sessionValue == nil {
		return false
	}
	return true
}

func AuthSessionMiddle() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		sessionValue := session.Get("uid")
		logrus.Println(sessionValue)
		if sessionValue == nil {
			c.Redirect(http.StatusFound, "/")
			return
		}

		uidInt, _ := strconv.Atoi(sessionValue.(string))

		if uidInt <= 0 {
			c.Redirect(http.StatusFound, "/")
			return
		}

		c.Set("uid", sessionValue)

		c.Next()
		return
	}
}
