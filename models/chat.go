package models

type Chat struct {
	Id        int   `json:"id"`
	UserId    int   `json:"user_id"`
	CreatedAt int64 `json:"created_at"`
	DeletedAt int64 `json:"deleted_at"`
}
