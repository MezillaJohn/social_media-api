package store

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
)

var ErrNotFound error = errors.New("record not found")

type Storage struct {
	Posts interface {
		Create(context.Context, *Post) error
		GetPostById(ctx context.Context, postId uuid.UUID) (*Post, error)
	}
	Users interface {
		Create(context.Context) error
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Posts: &PostStorage{db},
		Users: &UserStore{db},
	}
}
