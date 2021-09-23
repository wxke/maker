package config

import "github.com/go-pg/pg/v10"

func NewDB() *pg.DB {
	return pg.Connect(&pg.Options{
		Addr:     "localhost:5432",
		User:     "maker",
		Database: "maker",
		Password: "",
	})
}
