package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/reaperhero/go-gin-websocket/model"
	"github.com/sirupsen/logrus"
	"net/http"
)

func SaveAuthSession(c *gin.Context, user interface{}) { // info不能为struct
	session := sessions.Default(c)
	session.Set("uid", user)
	err := session.Save()
	if err != nil {
		logrus.WithField("[SaveAuthSession]:", err).Info(err)
	}
}

func GetSessionUserInfo(c *gin.Context) map[string]interface{} {
	uid := sessions.Default(c).Get("uid")
	user, _ := uid.(model.User)
	data := make(map[string]interface{})
	data["uid"] = user.ID
	data["username"] = user.Username
	data["avatar_id"] = user.AvatarId
	return data
}

func ClearAuthSession(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	if err := session.Save(); err != nil {
		logrus.WithField("ClearAuthSession", "失败").Error(err)
	}
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
		if sessionValue == nil {
			c.Redirect(http.StatusFound, "/")
			return
		}

		userInfo, ok := sessionValue.(model.User)
		if !ok {
			return
		}
		if userInfo.Username == "" || userInfo.Password == "" {
			c.Redirect(http.StatusFound, "/")
			return
		}

		c.Set("uid", sessionValue)

		c.Next()
		return
	}
}
