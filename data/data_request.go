package data

import (
	"database/sql"

	"github.com/dro14/yordamchi-api/models"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func (d *Data) CreateRequest(ctx *gin.Context, request *models.Request) error {
	query := "INSERT INTO requests (user_id, chat_id, started_at, finished_at, latency, chunks, errors, language, system_instruction, contents, response, finish_reason, model, cached_tokens, non_cached_tokens, tool_prompt_tokens, thought_tokens, response_tokens, price) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19)"
	var chunks sql.NullInt64
	if request.Chunks != 0 {
		chunks.Valid = true
		chunks.Int64 = request.Chunks
	}
	var errors sql.NullInt64
	if request.Errors != 0 {
		errors.Valid = true
		errors.Int64 = request.Errors
	}
	contents := make([]int64, len(request.Contents))
	for i, message := range request.Contents {
		contents[i] = message.Id
	}
	var cachedTokens sql.NullInt64
	if request.CachedTokens != 0 {
		cachedTokens.Valid = true
		cachedTokens.Int64 = request.CachedTokens
	}
	var toolPromptTokens sql.NullInt64
	if request.ToolPromptTokens != 0 {
		toolPromptTokens.Valid = true
		toolPromptTokens.Int64 = request.ToolPromptTokens
	}
	var thoughtTokens sql.NullInt64
	if request.ThoughtTokens != 0 {
		thoughtTokens.Valid = true
		thoughtTokens.Int64 = request.ThoughtTokens
	}
	args := []any{request.UserId, request.ChatId, request.StartedAt, request.FinishedAt, request.Latency, chunks, errors, request.Language, request.SystemInstruction, pq.Array(contents), request.Response.Id, request.FinishReason, request.Model, cachedTokens, request.NonCachedTokens, toolPromptTokens, thoughtTokens, request.ResponseTokens, request.Price}
	return d.dbExec(ctx, query, args...)
}
