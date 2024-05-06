package model

type ToDoCreateOrUpdateRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}
