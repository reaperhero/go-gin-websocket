package http

import (
	"github.com/gin-gonic/gin"
	"github.com/reaperhero/go-gin-websocket/middleware"
	"github.com/reaperhero/go-gin-websocket/model"
	"github.com/reaperhero/go-gin-websocket/utils"
	"net/http"
	"strconv"
)

func (h *handler) index(c *gin.Context) {
	userinfo := middleware.GetSessionUserInfo(c)
	if userinfo != nil {
		c.Redirect(http.StatusFound, "/home")
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
	u := struct {
		Username string `form:"username" binding:"required,max=16,min=2"`
		Password string `form:"password" binding:"required,max=32,min=6"`
		AvatarId string `form:"avatar_id" binding:"required,numeric"`
	}{}

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
	middleware.SaveAuthSession(c, user)
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
	c.HTML(http.StatusOK, "room.html", gin.H{
		"user_info": userInfo,
		"user_id":   2,
		"room_id":   roomId,
		"msg_list":  messageList,
	})
}

func (h *handler) wsHandler(c *gin.Context) {
	wsUpgrader.CheckOrigin = func(r *http.Request) bool { return true }
	wsContext, _ := wsUpgrader.Upgrade(c.Writer, c.Request, nil)
	defer wsContext.Close()

	go handleRead(wsContext)
	go handleWrite()

	select {}
}

func (h *handler) privateChat(c *gin.Context) {
	roomId := c.Query("room_id")
	toUid := c.Query("uid")
	userInfo := middleware.GetSessionUserInfo(c)
	uid := strconv.Itoa(int(userInfo["uid"].(uint)))
	msgList := h.usecase.GetLimitPrivateMsg(uid, toUid)
	c.HTML(http.StatusOK, "private_chat.html", gin.H{
		"user_info": userInfo,
		"msg_list":  msgList,
		"room_id":   roomId,
	})

}

func (h *handler) pagination(context *gin.Context) {

}
