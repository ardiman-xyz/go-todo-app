package web

type TodoResponse struct {
	Id     int    `json:"id"`
	Task   string `json:"task"`
	Status bool   `json:"status"`
}