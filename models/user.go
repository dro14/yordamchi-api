package models

type User struct {
	Id           int   `json:"id"`
	RegisteredAt int64 `json:"registered_at"`
}
