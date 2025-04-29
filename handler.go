package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

var todos []Todo
var nextID = 1

func todosHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getTodos(w, r)
	case http.MethodPost:
		createTodo(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func todoHandler(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/todos/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		getTodoByID(w, r, id)
	case http.MethodPut:
		updateTodo(w, r, id)
	case http.MethodDelete:
		deleteTodo(w, r, id)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get Query Parameters
	pageParam := r.URL.Query().Get("page")
	limitParam := r.URL.Query().Get("limit")

	page, err := strconv.Atoi(pageParam)
	if err != nil || page < 0 {
		page = 1
	}

	limit, err := strconv.Atoi(limitParam)
	if err != nil || limit < 0 {
		limit = 10
	}

	start := (page - 1) * limit
	end := start + limit
	total := len(todos)

	// Handle out of range
	if start > total {
		start = total
	}

	if end > total {
		end = total
	}

	pagedTodos := todos[start:end]

	response := BaseResponse{
		Success: true,
		Message: "Todos retrieved successfully",
		Data:    pagedTodos,
		Pagination: &Pagination{
			Page:  page,
			Limit: limit,
			Total: total,
		},
	}

	json.NewEncoder(w).Encode(response)
}

func getTodoByID(w http.ResponseWriter, r *http.Request, id int) {
	for _, todo := range todos {
		if todo.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(todo)
			return
		}
	}

	http.Error(w, "Todo not found", http.StatusOK)
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	if err := ValidateTodo(todo); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(BaseResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	todo.ID = nextID
	nextID++
	todos = append(todos, todo)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(BaseResponse{
		Success: true,
		Message: "Todo created successfully",
		Data:    todo,
	})
}

func updateTodo(w http.ResponseWriter, r *http.Request, id int) {
	var updatedTodo Todo
	if err := json.NewDecoder(r.Body).Decode(&updatedTodo); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	if err := ValidateTodo(updatedTodo); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(BaseResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	for i, todo := range todos {
		if todo.ID == id {
			todos[i].Title = updatedTodo.Title
			todos[i].Done = updatedTodo.Done
			json.NewEncoder(w).Encode(BaseResponse{
				Success: true,
				Message: "Todo updated successfully",
				Data:    todos[i],
			})
			return
		}
	}

	http.Error(w, "Todo not found", http.StatusNotFound)
}

func deleteTodo(w http.ResponseWriter, r *http.Request, id int) {
	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, "Todo not found", http.StatusNotFound)
}
