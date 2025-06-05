package models

type Message struct {
	Id        int64    `json:"id"`
	UserId    int64    `json:"user_id"`
	ChatId    int64    `json:"chat_id"`
	Role      string   `json:"role"`
	CreatedAt int64    `json:"created_at"`
	DeletedAt int64    `json:"deleted_at"`
	InReplyTo int64    `json:"in_reply_to"`
	Text      string   `json:"text"`
	Images    []string `json:"images"`
}
