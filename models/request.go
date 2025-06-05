package models

type Request struct {
	Id                int        `json:"id"`
	UserId            int        `json:"user_id"`
	ChatId            int        `json:"chat_id"`
	StartedAt         int64      `json:"started_at"`
	FinishedAt        int64      `json:"finished_at"`
	Latency           int64      `json:"latency"`
	Chunks            int        `json:"chunks"`
	Attempts          int        `json:"attempts"`
	Language          string     `json:"language"`
	SystemInstruction string     `json:"system_instruction"`
	Contents          []*Message `json:"contents"`
	Response          *Message   `json:"response"`
	StructuredOutput  string     `json:"structured_output"`
	ToolCalls         []string   `json:"tool_calls"`
	FinishReason      string     `json:"finish_reason"`
	Model             string     `json:"model"`
	PromptTokens      int        `json:"prompt_tokens"`
	ResponseTokens    int        `json:"response_tokens"`
	Price             float64    `json:"price"`
}
