package db

import (
	"context"
	"database/sql"
	"time"
)

func New(addr, maxIdleTime string, maxIdleConns, maxOpenConns int) (*sql.DB, error) {
	db, err := sql.Open("postgres", addr)
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(maxIdleConns)
	db.SetMaxOpenConns(maxOpenConns)

	duration, err := time.ParseDuration(maxIdleTime)

	if err != nil {
		return nil, err
	}

	db.SetConnMaxIdleTime(duration)

	ctx, cancle := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancle()

	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}

	return db, nil
}
