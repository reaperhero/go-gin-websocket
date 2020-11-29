package http

import (
	"github.com/gin-gonic/gin"
	"github.com/reaperhero/go-gin-websocket/middleware"
	"github.com/reaperhero/go-gin-websocket/model"
	"github.com/reaperhero/go-gin-websocket/utils"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (h *handler) index(c *gin.Context) {
	userinfo := middleware.GetSessionUserInfo(c)
	if userinfo != nil {
		c.Redirect(http.StatusFound, "/user/home")
		return
	}
	c.HTML(http.StatusOK, "login.html", gin.H{
		"OnlineUserCount": 10,
	})
	return
}

func (h *handler) home(c *gin.Context) {
	users := middleware.GetSessionUserInfo(c)
	rooms := []map[string]interface{}{
		{"id": 1, "num": 1},
		{"id": 2, "num": 2},
		{"id": 3, "num": 3},
		{"id": 4, "num": 4},
		{"id": 5, "num": 5},
		{"id": 6, "num": 6},
	}
	c.HTML(http.StatusOK, "index.html", gin.H{
		"rooms":     rooms,
		"user_info": users,
	})
	return
}

func (h *handler) userRegister(c *gin.Context) {
	var u model.User
	if err := c.ShouldBind(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 1001, "msg": err.Error()})
		return
	}
	if err := h.usecase.SaveUser(u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 1002,
			"msg":  "用户已经存在",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1001,
		"msg":  "注册成功",
	})
}

func (h *handler) userLogin(c *gin.Context) {
	var u model.User
	if err := c.ShouldBind(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 1001, "msg": err.Error()})
		return
	}
	user := h.usecase.FindUserByName(u.Username)
	if user.Username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"code": 1002, "msg": "用户不存在"})
		return
	}
	if !utils.CompareHashAndPassword(user.Password, u.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"code": 1003, "msg": "密码错误"})
		return
	}
	middleware.SaveAuthSession(c, u)
	c.JSON(http.StatusOK, gin.H{"code": 1000, "msg": "登陆成功"})
	return
}

func (h *handler) userLogout(c *gin.Context) {
	middleware.ClearAuthSession(c)
	c.Redirect(http.StatusFound, "/")
	return
}

func (h *handler) room(c *gin.Context) {
	roomId := c.Param("room_id")
	userInfo := middleware.GetSessionUserInfo(c)
	messageList := h.usecase.GetMessageByRoomId(roomId)
	for i, i2 := range messageList {
		logrus.Println(i, i2)
	}
	c.JSON(200, messageList)
	return
	c.HTML(http.StatusOK, "room.html", gin.H{
		"user_info":   userInfo,
		"room_id":     roomId,
		"messagelist": messageList,
	})
}
