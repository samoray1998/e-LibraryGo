package models

import "time"

type User struct {
	ID               int    `json:"id"`
	UserName         string `json:"userName"`
	Email            string `jsom:"email"`
	PasswordHash     string `json:"passwordHash"`
	CreatedAt        time.Time `json:"createdAt"`
	LastTimeSignedIn string `json:"lastTimeSignedIn"`
}
