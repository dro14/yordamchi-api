package models

type Message struct {
	Id        int      `json:"id"`
	UserId    int      `json:"user_id"`
	ChatId    int      `json:"chat_id"`
	Role      string   `json:"role"`
	CreatedAt int64    `json:"created_at"`
	DeletedAt int64    `json:"deleted_at"`
	InReplyTo int      `json:"in_reply_to"`
	Text      string   `json:"text"`
	Images    []string `json:"images"`
}
