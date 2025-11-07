package main

import (
	"fmt"
	"log"
	"net/http"
	"sample-inventory-go/handlers"
	"sample-inventory-go/middleware"
	"sample-inventory-go/store"

	"github.com/gorilla/mux"
)

func main() {
	inventoryStore := store.NewInventoryStore()
	apiHandler := handlers.NewAPIHandler(inventoryStore)

	router := mux.NewRouter()

	// Public routes
	router.HandleFunc("/login", apiHandler.Login).Methods("POST")

	// Authenticated routes
	authRouter := router.PathPrefix("/").Subrouter()
	authRouter.Use(func(next http.Handler) http.Handler {
		return middleware.AuthMiddleware(next, inventoryStore)
	})

	authRouter.HandleFunc("/logout", apiHandler.Logout).Methods("POST")
	authRouter.HandleFunc("/items", apiHandler.AddItem).Methods("POST")
	authRouter.HandleFunc("/items/{itemCode}", apiHandler.UpdateItem).Methods("PUT")
	authRouter.HandleFunc("/items/{itemCode}", apiHandler.DeleteItem).Methods("DELETE")
	authRouter.HandleFunc("/items", apiHandler.FetchItems).Methods("GET")

	port := ":8080"
	fmt.Printf("Server starting on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}
