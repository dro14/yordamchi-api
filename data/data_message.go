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
	query := "INSERT INTO messages (user_id, chat_id, role, created_at, in_reply_to, text, images, follow_ups, calls, responses) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id"
	var inReplyTo sql.NullInt64
	if message.InReplyTo != 0 {
		inReplyTo.Valid = true
		inReplyTo.Int64 = message.InReplyTo
	}
	var text sql.NullString
	if len(message.Text) > 0 {
		text.Valid = true
		text.String = message.Text
	}
	var calls sql.Null[[]byte]
	if len(message.Calls) > 0 {
		calls.Valid = true
		calls.V, _ = json.Marshal(message.Calls)
	}
	var responses sql.Null[[]byte]
	if len(message.Responses) > 0 {
		responses.Valid = true
		responses.V, _ = json.Marshal(message.Responses)
	}
	args := []any{message.UserId, message.ChatId, message.Role, message.CreatedAt, inReplyTo, text, pq.Array(message.Images), pq.Array(message.FollowUps), calls, responses}
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
