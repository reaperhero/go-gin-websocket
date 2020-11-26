package model

import "encoding/gob"

func init() {
	gob.Register(User{})
}

type User struct {
	Username string `form:"username"  binding:"required,max=16,min=2" db:"username"`
	Password string `form:"password"  binding:"required,max=32,min=6" db:"password"`
	AvatarId string `form:"avatar_id" binding:"required,numeric" db:"avatar_id"`
}
