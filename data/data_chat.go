package data

import (
	"context"

	"github.com/dro14/yordamchi-api/models"
	"github.com/lib/pq"
)

func (d *Data) CreateChat(ctx context.Context, chat *models.Chat) error {
	query := "INSERT INTO chats (user_id, created_at, updated_at) VALUES ($1, $2, $3) RETURNING id"
	args := []any{chat.UserId, chat.CreatedAt, chat.UpdatedAt}
	var id int64
	err := d.dbQueryRow(ctx, query, args, &id)
	if err != nil {
		return err
	}
	chat.Id = id
	return nil
}

func (d *Data) DeleteChats(ctx context.Context, chatIds []int64, deletedAt int64) error {
	query := "UPDATE chats SET deleted_at = $1 WHERE id = ANY($2)"
	args := []any{deletedAt, pq.Array(chatIds)}
	return d.dbExec(ctx, query, args...)
}
