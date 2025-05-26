package main

import (
	"log"

	"github.com/MezillaJohn/social_media-api/internal/db"
	"github.com/MezillaJohn/social_media-api/internal/env"
	"github.com/MezillaJohn/social_media-api/internal/store"
)

func main() {
	cfg := config{
		addr: env.GetStringEnv("ADDR", ":8080"),
		db: dbConfig{
			addr:         env.GetStringEnv("DB_ADDR", "postgres://postgres:admin@localhost/social?sslmode=disable"),
			maxOpenConns: env.GetIntEnv("MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetIntEnv("MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetStringEnv("MAX_IDLE_TIME", "15m"),
		},
		env: env.GetStringEnv("ENV", "development"),
	}

	db, err := db.New(
		cfg.db.addr,
		cfg.db.maxIdleTime,
		cfg.db.maxIdleConns,
		cfg.db.maxOpenConns,
	)

	if err != nil {
		log.Panic(err)
	}

	log.Printf("database connection pool established")

	defer db.Close()

	store := store.NewStorage(db)

	app := application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
