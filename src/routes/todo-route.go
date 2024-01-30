package routes

import (
	"github.com/vanditkhurana/todo_api/src/controllers"
	"github.com/gorilla/mux"
)

type WelcomeResponse struct {
	Message string `json:"message"`
}

func RegisterRoutes(router *mux.Router) {

	// Defining API endpoints
	router.HandleFunc("/", controllers.WelcomeHandler)
	router.HandleFunc("/todos", controllers.GetAllTodos).
	Queries(
		"page", "{page:[a-zA-Z0-9]+}",
		"limit", "{limit:[a-zA-Z0-9]+}",
		"status", "{status:[a-zA-Z0-9]+}").Methods("GET")
	router.HandleFunc("/todos/{user_id}", controllers.GetTodo).Methods("GET")
	router.HandleFunc("/todos", controllers.CreateTodo).Methods("POST")
	router.HandleFunc("/todos/{user_id}", controllers.UpdateTodo).Methods("PUT")
	router.HandleFunc("/todos/{user_id}", controllers.DeleteTodo).Methods("DELETE")


}


