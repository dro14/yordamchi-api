package models

import "google.golang.org/genai"

type Message struct {
	Id        int64                     `json:"id,omitempty"`
	UserId    int64                     `json:"user_id,omitempty"`
	ChatId    int64                     `json:"chat_id,omitempty"`
	Role      string                    `json:"role,omitempty"`
	CreatedAt int64                     `json:"created_at,omitempty"`
	DeletedAt int64                     `json:"deleted_at,omitempty"`
	InReplyTo int64                     `json:"in_reply_to,omitempty"`
	Text      string                    `json:"text,omitempty"`
	Images    []string                  `json:"images,omitempty"`
	FollowUps []string                  `json:"follow_ups,omitempty"`
	Calls     []*genai.FunctionCall     `json:"calls,omitempty"`
	Responses []*genai.FunctionResponse `json:"responses,omitempty"`
}
