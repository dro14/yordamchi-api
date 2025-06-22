package models

type Chat struct {
	Id        int64  `json:"id,omitempty"`
	UserId    int64  `json:"user_id,omitempty"`
	CreatedAt int64  `json:"created_at,omitempty"`
	UpdatedAt int64  `json:"updated_at,omitempty"`
	DeletedAt int64  `json:"deleted_at,omitempty"`
	Name      string `json:"name,omitempty"`
}
