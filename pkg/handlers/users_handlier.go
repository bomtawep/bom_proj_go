package handlers

import (
	"log"
	"net/http"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	log.Println(w, "List of products")
}
