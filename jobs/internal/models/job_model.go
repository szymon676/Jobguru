package models

type Job struct {
	ID          int64    `json:"id"`
	Title       string   `json:"title"`
	Skills      []string `json:"skills"`
	Category    string   `json:"category"`
	Description string   `json:"description"`
}

type BindJob struct {
	Title       string   `json:"title"`
	Skills      []string `json:"skills"`
	Category    string   `json:"category"`
	Description string   `json:"description"`
}
