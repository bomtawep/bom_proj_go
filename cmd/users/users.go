package users

import (
	"fmt"
	"gorm.io/gorm"
	"net/http"
)

type User struct {
	gorm.Model
	Name  string
	Email string `gorm:"unique"`
}

var (
	db *gorm.DB
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	user := User{Name: "John Doe", Email: "john@example.com"}

	result := db.Create(&user)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "User created successfully! ID: %d", user.ID)
}
