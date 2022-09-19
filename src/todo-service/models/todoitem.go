package models

type TodoItem struct {
	Text   string `json:"text"`
	Status string `json:"status"`
	Tags   string `json:"tags"`
	Color  string `json:"color"`
}
