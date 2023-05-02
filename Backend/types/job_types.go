package types

import "time"

type Job struct {
	ID          uint64    `json:"id"`
	UserID      uint      `json:"user_id"`
	Title       string    `json:"title"`
	Company     string    `json:"company"`
	Skills      []string  `json:"skills"`
	Salary      int       `json:"salary"`
	Description string    `json:"description"`
	Currency    string    `json:"currency"`
	Date        time.Time `json:"date"`
	Location    string    `json:"location"`
}

type JobReq struct {
	UserID      uint     `json:"user_id"`
	Title       string   `json:"title"`
	Company     string   `json:"company"`
	Skills      []string `json:"skills"`
	Salary      int      `json:"salary"`
	Description string   `json:"description"`
	Currency    string   `json:"currency"`
	Date        string   `json:"date"`
	Location    string   `json:"location"`
}
