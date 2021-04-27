package entity

import "time"

type DemoEntity struct {
	Id        int64
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
