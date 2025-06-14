package data

import (
	"database/sql"
	"encoding/json"

	"github.com/dro14/yordamchi-api/models"
	"github.com/dro14/yordamchi-api/utils/f"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func (d *Data) CreateMessage(ctx *gin.Context, message *models.Message) error {
	query := "INSERT INTO messages (user_id, chat_id, role, created_at, in_reply_to, text, images) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id"
	var nullInReplyTo sql.NullInt64
	if message.InReplyTo != 0 {
		nullInReplyTo.Valid = true
		nullInReplyTo.Int64 = message.InReplyTo
	}
	var nullText sql.NullString
	if len(message.Text) > 0 {
		nullText.Valid = true
		nullText.String = message.Text
	}
	var functionCalls []byte
	if len(message.FunctionCalls) > 0 {
		functionCalls, _ = json.Marshal(message.FunctionCalls)
	}
	var functionResponses []byte
	if len(message.FunctionResponses) > 0 {
		functionResponses, _ = json.Marshal(message.FunctionResponses)
	}
	var structuredOutput []byte
	if len(message.StructuredOutput) > 0 {
		structuredOutput = []byte(message.StructuredOutput)
	}
	args := []any{message.UserId, message.ChatId, message.Role, message.CreatedAt, nullInReplyTo, nullText, pq.Array(message.Images), functionCalls, functionResponses, structuredOutput}
	var id int64
	err := d.dbQueryRow(ctx, query, args, &id)
	if err != nil {
		return err
	}
	message.Id = id
	query = "UPDATE chats SET updated_at = $1 WHERE id = $2"
	args = []any{message.CreatedAt, message.ChatId}
	return d.dbExec(ctx, query, args...)
}

func (d *Data) DeleteUntil(ctx *gin.Context, chatId int64, messageId int64) error {
	query := "UPDATE messages SET deleted_at = $1 WHERE chat_id = $2 AND id >= $3 AND deleted_at IS NULL"
	args := []any{f.Now(), chatId, messageId}
	return d.dbExec(ctx, query, args...)
}
