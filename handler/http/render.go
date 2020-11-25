package http

import (
	"github.com/gin-gonic/gin"
	"github.com/reaperhero/go-gin-websocket/model"
	"net/http"
)

func (h *handler) test(context *gin.Context) {

}

func (h *handler) index(c *gin.Context) {

	c.HTML(http.StatusOK, "index.html", gin.H{})
}

func (h *handler) userRegister(c *gin.Context) {
	var u model.User
	if err := c.ShouldBind(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.HTML(http.StatusOK, "index.html", gin.H{})
}
