package models

type Request struct {
	Id                int64      `json:"id"`
	UserId            int64      `json:"user_id"`
	ChatId            int64      `json:"chat_id"`
	StartedAt         int64      `json:"started_at"`
	FinishedAt        int64      `json:"finished_at"`
	Latency           int64      `json:"latency"`
	Chunks            int64      `json:"chunks"`
	Attempts          int64      `json:"attempts"`
	Language          string     `json:"language"`
	SystemInstruction string     `json:"system_instruction"`
	Contents          []*Message `json:"contents"`
	Response          *Message   `json:"response"`
	FinishReason      string     `json:"finish_reason"`
	Model             string     `json:"model"`
	PromptTokens      int64      `json:"prompt_tokens"`
	ResponseTokens    int64      `json:"response_tokens"`
	Price             float64    `json:"price"`
}
