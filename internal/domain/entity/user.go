package entity

import "time"

type UserEntity struct {
	UserID    int64
	Mobile    string
	Password  string
	Nickname  string
	Avatar    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
