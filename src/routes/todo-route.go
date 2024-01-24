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
	router.HandleFunc("/todos?page={page}&limit={limit}&status={status}", controllers.GetAllTodos).Methods("GET")
	router.HandleFunc("/todos/{user_id}", controllers.GetTodo).Methods("GET")
	router.HandleFunc("/todos", controllers.CreateTodo).Methods("POST")
	router.HandleFunc("/todos/{user_id}", controllers.UpdateTodo).Methods("PUT")
	router.HandleFunc("/todos/{user_id}", controllers.DeleteTodo).Methods("DELETE")


}


