package handlers

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

// KeyProduct is a key used for the Category object in the context
type KeyProduct struct{}

type Categories struct{
	l *log.Logger
}

func NewCategories(l *log.Logger) *Categories{
	return &Categories{l}
}

func getCategoryID(r *http.Request) int {
	// parse the category id from the url
	vars := mux.Vars(r)

	// convert the id into an integer and return
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		// should never happen
		panic(err)
	}
	return id
}
