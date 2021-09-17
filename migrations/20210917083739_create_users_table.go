package main

import (
	"github.com/go-pg/pg/v10/orm"
	migrations "github.com/robinjoseph08/go-pg-migrations/v3"
)

func init() {
	up := func(db orm.DB) error {
		_, err := db.Exec(`
			CREATE TABLE users (
				id UUID PRIMARY KEY UNIQUE DEFAULT uuid_generate_v4(),
				name TEXT NOT NULL UNIQUE,
				email TEXT NOT NULL UNIQUE,
				password TEXT NOT NULL,
				created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
				updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
			);
			CREATE INDEX user_id_index ON users(id);
			CREATE INDEX user_name_index ON users(name);
			CREATE INDEX user_created_at_index ON users(created_at DESC NULLS LAST);
		`)
		return err
	}

	down := func(db orm.DB) error {
		_, err := db.Exec("DROP TABLE users")
		return err
	}

	opts := migrations.MigrationOptions{}

	migrations.Register("20210917083739_create_users_table", up, down, opts)
}
