package models

type Chat struct {
	Id        int64  `json:"id"`
	UserId    int64  `json:"user_id"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
	DeletedAt int64  `json:"deleted_at"`
	Name      string `json:"name"`
}
