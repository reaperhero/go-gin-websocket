package model

import (
	"time"
)

type Message struct {
	ID        uint
	UserId    int
	ToUserId  int
	RoomId    int
	Content   string
	ImageUrl  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
