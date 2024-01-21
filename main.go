package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vanditkhurana/todo_api/src/db"
	"github.com/vanditkhurana/todo_api/src/routes"
	"github.com/vanditkhurana/todo_api/src/controllers"
)

func main() {
	session := db.InitScyllaDB()
	defer session.Close()

	// Initialize routes from routers.go
	router := mux.NewRouter()
	controllers.Initialize(session)
	routes.RegisterRoutes(router)

	log.Fatal(http.ListenAndServe(":8080", router))
}
