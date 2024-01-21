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

func getAllTodos(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	status := r.URL.Query().Get("status")
	sortBy := r.URL.Query().Get("sort")
}

func getTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	// Parse UUID
	id, err := gocql.ParseUUID(idStr)
	if err != nil {
		http.Error(w, "Invalid UUID", http.StatusBadRequest)
		return
	}

	// Retrieve TODO item from the database
	var todo models.Todo
	if err := session.Query(`
		SELECT id, user_id, title, description, status, created, updated
		FROM todos
		WHERE user_id = ?
	`, id).Scan(&todo.ID, &todo.UserID, &todo.Title, &todo.Description, &todo.Status, &todo.Created, &todo.Updated); err != nil {
		http.Error(w, "TODO item not found", http.StatusNotFound)
		return
	}

	// Respond with the retrieved TODO item
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}

func createTodo(w http.ResponseWriter, r *http.Request) {
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
	if err := session.Query(`
		INSERT INTO todos (id, user_id, title, description, status, created, updated)
		VALUES (?, ?, ?, ?, ?, ?, ?)
	`, todo.ID, todo.UserID, todo.Title, todo.Description, todo.Status, todo.Created, todo.Updated).Exec(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with the created TODO item
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todo)
}

func updateTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	// Parse UUID
	id, err := gocql.ParseUUID(idStr)
	if err != nil {
		http.Error(w, "Invalid UUID", http.StatusBadRequest)
		return
	}

	var todo models.Todo
	// Parse request body
	err = json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Update the TODO item
	// Placeholder logic
	todo.ID = id
	todo.Updated = time.Now()

	// Implement logic to update a TODO item in the database by ID
	// Placeholder logic
	if err := session.Query(`
		UPDATE todos
		SET user_id = ?, title = ?, description = ?, status = ?, updated = ?
		WHERE id = ?
	`, todo.UserID, todo.Title, todo.Description, todo.Status, todo.Updated, todo.ID).Exec(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with the updated TODO item
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	// Parse UUID
	id, err := gocql.ParseUUID(idStr)
	if err != nil {
		http.Error(w, "Invalid UUID", http.StatusBadRequest)
		return
	}

	// Delete TODO item from the database
	if err := session.Query(`
		DELETE FROM todos
		WHERE id = ?
	`, id).Exec(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with success message
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "TODO item deleted successfully")
}