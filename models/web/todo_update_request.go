package web

type TodoUpdateRequest struct {
	Id int `validate:"required" json:"id"`
	Task string `validate:"required" json:"task"`
	Status bool `json:"status"`
}