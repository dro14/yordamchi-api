package data

import (
	"database/sql"

	"github.com/dro14/yordamchi-api/models"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func (d *Data) CreateRequest(ctx *gin.Context, request *models.Request) error {
	query := "INSERT INTO requests (user_id, chat_id, started_at, finished_at, latency, chunks, attempts, language, system_instruction, contents, response, structured_output, tool_calls, finish_reason, model, prompt_tokens, response_tokens, price) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18)"
	contents := make([]int64, len(request.Contents))
	for i, message := range request.Contents {
		contents[i] = message.Id
	}
	var nullResponse sql.NullInt64
	var nullStructuredOutput sql.NullString
	if request.Response != nil {
		nullResponse.Valid = true
		nullResponse.Int64 = request.Response.Id
	}
	if request.StructuredOutput != "" {
		nullStructuredOutput.Valid = true
		nullStructuredOutput.String = request.StructuredOutput
	}
	args := []any{request.UserId, request.ChatId, request.StartedAt, request.FinishedAt, request.Latency, request.Chunks, request.Attempts, request.Language, request.SystemInstruction, pq.Array(contents), nullResponse, nullStructuredOutput, pq.Array(request.ToolCalls), request.FinishReason, request.Model, request.PromptTokens, request.ResponseTokens, request.Price}
	return d.dbExec(ctx, query, args...)
}
