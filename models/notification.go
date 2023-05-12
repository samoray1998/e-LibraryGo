package models

import "time"

type Notification struct {
	ID        int
	User      *User
	Message   string
	Metadata  map[string]interface{}
	CreatedAt time.Time
}
