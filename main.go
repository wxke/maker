package main

import (
	"fmt"
	"maker/config"
	"maker/models"
	"maker/routers"
	"net/http"
	"time"
)

func main() {
	router := routers.InitRouter()
	db := config.GetConnection()
	defer db.Close()
	db.QueryOne(&models.User{}, `
		INSERT INTO users (name, email) VALUES (?name, ?email, ?password)
		RETURNING id
	`, &models.User{Name: "wxk", Email: "wxk", Password: "wxk"})
	fmt.Println(db.Query(&models.User{}, `SELECT * FROM users`))
	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
