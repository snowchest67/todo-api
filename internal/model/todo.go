package model

type CreateTodoRequest struct {
	Title string `json:"title"`
}

type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}