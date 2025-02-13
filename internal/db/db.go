package db

import (
	"context"

	"gitlab.aaronhess.xyz/viddler/viddler-blog-api/internal/db/ent"

	_ "gitlab.aaronhess.xyz/viddler/viddler-blog-api/internal/db/sqlite3"
)

type DBConfig struct {
	DSN string
}

type DB struct {
	*ent.Client
}

func Init(ctx context.Context, params *DBConfig) (*DB, error) {
	client, err := ent.Open("sqlite3", params.DSN)
	if err != nil {
		return nil, err
	}
	db := &DB{Client: client}
	if err := db.Migrate(ctx); err != nil {
		return nil, err
	}
	return db, nil
}

func (db *DB) Migrate(ctx context.Context) error {
	return db.Schema.Create(ctx)
}
