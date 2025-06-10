package data

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/dro14/yordamchi-api/utils/e"
)

const retryAttempts = 5

func (d *Data) dbExec(ctx context.Context, query string, args ...any) error {
	var lastErr error
	for range retryAttempts {
		_, err := d.db.ExecContext(ctx, query, args...)
		if err == nil {
			return nil
		}
		lastErr = err
		log.Printf("can't exec: %s\n%s", err, query)
		if errors.Is(err, sql.ErrNoRows) {
			return e.ErrNotFound
		}
	}
	log.Printf("failed to exec after %d attempts", retryAttempts)
	return lastErr
}

func (d *Data) dbQueryRow(ctx context.Context, query string, args []any, dest ...any) error {
	var lastErr error
	for range retryAttempts {
		err := d.db.QueryRowContext(ctx, query, args...).Scan(dest...)
		if err == nil {
			return nil
		}
		lastErr = err
		log.Printf("can't query row: %s\n%s", err, query)
		if errors.Is(err, sql.ErrNoRows) {
			return e.ErrNotFound
		}
	}
	log.Printf("failed to query row after %d attempts", retryAttempts)
	return lastErr
}
