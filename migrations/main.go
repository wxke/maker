package main

import (
	"log"
	"maker/config"
	"maker/models"
	"os"

	migrations "github.com/robinjoseph08/go-pg-migrations/v3"
)

const directory = "migrations"

func main() {
	db := config.NewDB()

	err := migrations.Run(db, directory, os.Args)
	if err != nil {
		log.Fatalln(err)
	}

	models.Close()
}
