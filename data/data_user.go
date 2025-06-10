package data

import "context"

func (d *Data) CreateUser(ctx context.Context, registeredAt int64) (int64, error) {
	query := "INSERT INTO users (registered_at) VALUES ($1) RETURNING id"
	args := []any{registeredAt}
	var id int64
	err := d.dbQueryRow(ctx, query, args, &id)
	return id, err
}
