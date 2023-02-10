package main

import (
	"github.com/Schariss/category-api/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"time"
)


func main(){
	l := log.New(os.Stdout, "category-api", log.LstdFlags)
	ch := handlers.NewCategories(l)

	sm := mux.NewRouter()

	getR := sm.Methods(http.MethodGet).Subrouter()
	getR.HandleFunc("/categories", ch.GetList)
	getR.HandleFunc("/categories/{id:[0-9]+}", ch.GetSingle)

	postR := sm.Methods(http.MethodPost).Subrouter()
	postR.HandleFunc("/categories/", ch.Create)
	postR.Use(ch.MiddlewareValidateCategory)

	sm.HandleFunc("/",  func(rw http.ResponseWriter, r *http.Request){
		rw.Write([]byte("<h1 style='text-align:center;'>Welcome to Home page</h1>"))
	})

	server := &http.Server{
		Addr: ":9091",
		Handler: sm,
		IdleTimeout: 120*time.Second,
		ReadTimeout: 1*time.Second,
		WriteTimeout: 1*time.Second,
	}

	l.Fatal(server.ListenAndServe())
}
