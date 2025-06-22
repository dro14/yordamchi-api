package models

type Request struct {
	Id                int64      `json:"id,omitempty"`
	UserId            int64      `json:"user_id,omitempty"`
	ChatId            int64      `json:"chat_id,omitempty"`
	StartedAt         int64      `json:"started_at,omitempty"`
	FinishedAt        int64      `json:"finished_at,omitempty"`
	Latency           int64      `json:"latency,omitempty"`
	Chunks            int64      `json:"chunks,omitempty"`
	Errors            int64      `json:"errors,omitempty"`
	Language          string     `json:"language,omitempty"`
	SystemInstruction string     `json:"system_instruction,omitempty"`
	Contents          []*Message `json:"contents,omitempty"`
	Response          *Message   `json:"response,omitempty"`
	FinishReason      string     `json:"finish_reason,omitempty"`
	Model             string     `json:"model,omitempty"`
	CachedTokens      int64      `json:"cached_tokens,omitempty"`
	NonCachedTokens   int64      `json:"non_cached_tokens,omitempty"`
	ToolPromptTokens  int64      `json:"tool_prompt_tokens,omitempty"`
	ThoughtTokens     int64      `json:"thought_tokens,omitempty"`
	ResponseTokens    int64      `json:"response_tokens,omitempty"`
	Price             float64    `json:"price,omitempty"`
}
