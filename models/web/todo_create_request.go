package web

type TodoCreateRequest struct {
	Task string `validate:"required" json:"task"`
}