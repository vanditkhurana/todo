package routes

import (
	"fmt"
	"log"
	"net/http"
	"github.com/vanditkhurana/todo_api/src/controllers"
	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router) {

	// Defining API endpoints
	router.HandleFunc("/todos", controllers.getAllTodos).Methods("GET")
	router.HandleFunc("/todos/{id}", controllers.getTodo).Methods("GET")
	router.HandleFunc("/todos", controllers.createTodo).Methods("POST")
	router.HandleFunc("/todos/{id}", controllers.updateTodo).Methods("PUT")
	router.HandleFunc("/todos/{id}", controllers.deleteTodo).Methods("DELETE")

	// Starting the server
	port := 8080
	fmt.Printf("Server is running on :%d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}
