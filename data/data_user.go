package data

import "github.com/gin-gonic/gin"

func (d *Data) CreateUser(ctx *gin.Context, registeredAt int64) (int64, error) {
	query := "INSERT INTO users (registered_at) VALUES ($1) RETURNING id"
	args := []any{registeredAt}
	var id int64
	err := d.dbQueryRow(ctx, query, args, &id)
	return id, err
}
