package models

type User struct {
	Id           int64 `json:"id,omitempty"`
	RegisteredAt int64 `json:"registered_at,omitempty"`
}
