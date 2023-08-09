package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kisahtegar/go-jwt-mux/controllers/authController"
	"github.com/kisahtegar/go-jwt-mux/controllers/productController"
	"github.com/kisahtegar/go-jwt-mux/middlewares"
	"github.com/kisahtegar/go-jwt-mux/models"
)

func main() {
	// Connect to the database
	models.ConnectDatabase()

	r := mux.NewRouter()

	r.HandleFunc("/login", authController.Login).Methods("POST")
	r.HandleFunc("/register", authController.Register).Methods("POST")
	r.HandleFunc("/logout", authController.Logout).Methods("GET")

	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/products", productController.Index).Methods("GET")
	api.Use(middlewares.JWTMiddleware)

	log.Fatal(http.ListenAndServe(":8080", r))
}
