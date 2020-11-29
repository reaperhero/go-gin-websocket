package model

import (
	"encoding/gob"
	"time"
)

func init() {
	gob.Register(User{})
}

type User struct {
	ID        uint      `db:"id"`
	Username  string    `form:"username"  binding:"required,max=16,min=2" db:"username"`
	Password  string    `form:"password"  binding:"required,max=32,min=6" db:"password"`
	AvatarId  string    `form:"avatar_id" binding:"required,numeric" db:"avatar_id"`
	CreatedAt time.Time `time_format:"2006-01-02 15:04:05" db:"created_at"`
	UpdatedAt time.Time `time_format:"2006-01-02 15:04:05" db:"updated_at"`
	DeletedAt time.Time `time_format:"2006-01-02 15:04:05" db:"deleted_at"`
}
