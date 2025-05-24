package store

import (
	"context"
	"database/sql"
)

type UserStore struct {
	db *sql.DB
}

func (user *UserStore) Create(ctx context.Context) error {
	return nil
}
