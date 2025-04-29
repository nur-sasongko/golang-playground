package main

import "errors"

type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

type BaseResponse struct {
	Success    bool        `json:"success"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data",omitempty`
	Pagination *Pagination `json:"pagination,omitempty"`
}

type Pagination struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
	Total int `json:"total"`
}

func ValidateTodo(todo Todo) error {
	if todo.Title == "" {
		return errors.New("title is required")
	}
	if len(todo.Title) < 3 {
		return errors.New("title must be at least 3 characters")
	}
	return nil
}
