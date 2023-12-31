package user

import (
	"pustaka-api/book"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email string `gorm:"unique"`
	Password string

	//Relationship
	Books []book.Book
}