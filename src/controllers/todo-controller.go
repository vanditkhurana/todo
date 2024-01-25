package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gocql/gocql"
	"github.com/gorilla/mux"
	"github.com/vanditkhurana/todo_api/src/db"
	"github.com/vanditkhurana/todo_api/src/models"
)

func GetAllTodos(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()

	page := params.Get("page")
	limit := params.Get("limit")
	status := params.Get("status")
	// Default values
	if page == "" {
		page = "1"
	}
	if limit == "" {
		limit = "10"
	}

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		http.Error(w, "Invalid page value", http.StatusBadRequest)
		return
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		http.Error(w, "Invalid limit value", http.StatusBadRequest)
		return
	}

	// Calculate offset
	offset := (pageInt - 1) * limitInt
	fmt.Printf("%d", offset)
	// Initial query
	query := "SELECT id, user_id, title, description, status, created, updated FROM todos_new"
	
	// Appending query filter on the basis of status
	if status != "" {
		query += " WHERE status = ?"
	}

	// Sort by created 
	query += " ORDER BY created DESC"

	// Query for pagination
	query += " LIMIT ?"
	fmt.Printf(query)
	iter := db.Session.Query(query, status, limitInt).Iter()

	var todos []models.Todo
	var todo models.Todo

	for iter.Scan(&todo.ID, &todo.User_ID, &todo.Title, &todo.Description, &todo.Status, &todo.Created, &todo.Updated) {
		todos = append(todos, todo)
	}

	if err := iter.Close(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return the result as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

func GetTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["user_id"]

	// Retrieve TODO item from the database
	var todo models.Todo
	if err := db.Session.Query(`
		SELECT id, user_id, title, description, status, created, updated
		FROM todos_new
		WHERE user_id = ?`, userId).Scan(&todo.ID, &todo.User_ID, &todo.Title, &todo.Description, &todo.Status, &todo.Created, &todo.Updated); err != nil {
		http.Error(w, "TODO item not found", http.StatusNotFound)
		return
	}

	// Respond with the retrieved TODO item
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Set default values
	todo.ID = gocql.TimeUUID()
	todo.Created = time.Now()
	todo.Updated = time.Now()

	// Implement logic to create a new TODO item in the database
	// Placeholder logic
	if err := db.Session.Query(`
		INSERT INTO todos (id, user_id, title, description, status, created, updated)
		VALUES (?, ?, ?, ?, ?, ?, ?)
	`, &todo.ID, &todo.User_ID, &todo.Title, &todo.Description, &todo.Status, &todo.Created, &todo.Updated).Exec(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with the created TODO item
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todo)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["user_id"]

	var todo models.Todo
	// Parse request body
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Update the TODO item
	// Placeholder logic
	todo.User_ID = userID
	todo.Updated = time.Now()

	// Implement logic to update a TODO item in the database by ID
	// Placeholder logic
	if err := db.Session.Query(`
		UPDATE todos
		SET title = ?, description = ?, status = ?, updated = ?
		WHERE user_id = ?
	`, &todo.Title, &todo.Description, &todo.Status, &todo.Updated, &todo.User_ID).Exec(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with the updated TODO item
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["user_id"]

	// Delete TODO item from the database
	if err := db.Session.Query(`
		DELETE FROM todos
		WHERE user_id = ?
	`, userID).Exec(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with success message
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "TODO item deleted successfully!")
}

func WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	response := models.WelcomeResponse{"Welcome to TODOS API"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}