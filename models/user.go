package models

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"users"`
	ID            uuid.UUID `bun:"type:uuid,default:gen_random_uuid(),pk,notnull,unique:group_name"`
	Name          string    `bun:"name,notnull"`
	Email         string    `bun:"email,notnull,unique"`
	password      string    `bun:"password,notnull"`
	CreatedAt     time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt     time.Time `bun:",nullzero,notnull,default:current_timestamp"`
}

func (*User) BeforeCreateTableQuery(ctx context.Context, query *bun.CreateTableQuery) error {
	return nil
}

func (*User) AfterCreateTableQuery(ctx context.Context, query *bun.CreateTableQuery) error {
	_, err := query.DB().NewCreateIndex().
		Model((*User)(nil)).
		Index("name_idx").
		Column("name").
		Exec(ctx)
	if err != nil {
		return err
	}

	// 希望是倒叙 但是还没找到 倒叙的方法
	_, err = query.DB().NewCreateIndex().
		Model((*User)(nil)).
		Index("created_at_idx").
		Column("created_at").
		Exec(ctx)

	return err
}
