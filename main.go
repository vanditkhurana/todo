package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/vanditkhurana/todo_api/src/db"
	"github.com/vanditkhurana/todo_api/src/routes"
)

func main() {
	session := db.Session
	defer session.Close()

	router := mux.NewRouter()
	routes.RegisterRoutes(router)

	// Starting the server
	port := 8080
	fmt.Printf("Server is running on :%d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}

