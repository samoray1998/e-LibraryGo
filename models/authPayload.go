package models

type AuthPayload struct {
	Token string `json:"token"`
	User  *User  `json:user`
}
