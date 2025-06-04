package models

type Chat struct {
	Id        int `json:"id"`
	UserId    int `json:"user_id"`
	CreatedAt int `json:"created_at"`
	DeletedAt int `json:"deleted_at"`
}
