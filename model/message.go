package model

import (
	"time"
)

type Message struct {
	ID        uint      `db:"id"`
	UserId    int       `db:"user_id"`
	ToUserId  int       `db:"to_user_id"`
	RoomId    int       `db:"room_id"`
	Content   string    `db:"content"`
	ImageUrl  string    `db:"image_url"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	DeletedAt time.Time `db:"deleted_at"`
}

type PrivateMessage struct {
}
