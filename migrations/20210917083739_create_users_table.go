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
				name VARCHAR(64) NOT NULL UNIQUE,
				email VARCHAR(64) NOT NULL UNIQUE,
				password VARCHAR(64) NOT NULL,
				created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
				updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
			);
			CREATE INDEX user_created_at_index ON users USING btree (created_at DESC);
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
