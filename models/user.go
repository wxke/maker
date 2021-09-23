package models

import (
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `form:"id" json:"ID" pg:",notnull"`
	Name      string    `form:"name" json:"Name" binding:"required" pg:",notnull"`
	Email     string    `form:"email" json:"Email" pg:",notnull"`
	Password  string    `form:"password" json:"Password" pg:",notnull"`
	CreatedAt time.Time `form:"created_at" json:"CreatedAt" pg:",notnull"`
	UpdatedAt time.Time `form:"updated_at" json:"UpdatedAt" pg:",notnull"`
}

func (user *User) FindByName() error {
	err := db.Model(user).Where("? = ?", pg.Ident("name"), user.Name).Select()
	return err
}
