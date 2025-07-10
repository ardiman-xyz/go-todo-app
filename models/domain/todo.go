package domain

type TodoList struct {
	ID     int    `json:"id"`
	Task   string `json:"task"`
	Status bool   `json:"status"`
}