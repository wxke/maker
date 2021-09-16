package main

import (
	"context"
	"fmt"
	"maker/config"
	"maker/models"
	"maker/routers"
	"net/http"
	"time"
)

func main() {
	router := routers.InitRouter()

	db := config.NewDB()
	ctx := context.Background()

	if _, err := db.NewCreateTable().Model((*models.User)(nil)).Exec(ctx); err != nil {
		panic(err)
	}
	values := map[string]interface{}{
		"name":     "111",
		"email":    "111",
		"password": "111",
	}
	r, err := db.NewInsert().Model(&values).TableExpr("users").Exec(ctx)
	fmt.Println(r, err)

	ms := make([]map[string]interface{}, 0)
	err = db.NewSelect().
		Model(&models.User{}).
		Scan(ctx, &ms)
	fmt.Println(r, err, ms)

	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
