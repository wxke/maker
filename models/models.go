package models

import (
	"context"
	"fmt"
	"maker/config"

	"github.com/go-pg/pg/v10"
)

var db *pg.DB

type dbLogger struct{}

func init() {
	db = config.NewDB()
	db.AddQueryHook(dbLogger{})
}

func Close() {
	defer db.Close()
}

func (d dbLogger) BeforeQuery(c context.Context, q *pg.QueryEvent) (context.Context, error) {
	return c, nil
}

func (d dbLogger) AfterQuery(c context.Context, q *pg.QueryEvent) error {
	sql, _ := q.FormattedQuery()
	fmt.Println(string(sql))
	return nil
}
