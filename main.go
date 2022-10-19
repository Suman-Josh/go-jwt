package main

import (
	"fmt"
	"log"
	"net/http"

	"jwt-task/connection"
	"jwt-task/models"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func initializeRouter() {
	r := mux.NewRouter()
	r.HandleFunc("/login", Login).Methods("POST")
	r.HandleFunc("/refresh", Refresh).Methods("GET")
	r.HandleFunc("/users", GetUser).Methods("GET")
	r.HandleFunc("/users/{id}", GetUser).Methods("GET")
	r.HandleFunc("/users", CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))
}

func main() {
	DB, err = connection.SetupDB()
	if err != nil {
		log.Panic(err)
		return
	}
	fmt.Println("Connected to database")

	DB.AutoMigrate(&models.User{})
	initializeRouter()
}
