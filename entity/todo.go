package entity

// Todo entity has title and description
type Todo struct {
	Title  string `json:"title"`
	IsDone bool   `json:"isDone"`
}