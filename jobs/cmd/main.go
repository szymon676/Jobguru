package main

import (
	"fmt"

	"github.com/szymon676/job-guru/jobs/internal/database"
	"github.com/szymon676/job-guru/jobs/internal/handlers"
)

func main() {
	dsn := "host=localhost port=5432 user=postgres password=1234 dbname=jobguru sslmode=disable"
	db, err := database.NewDatabase("postgres", dsn)

	if err != nil {
		fmt.Printf("Error creating database")
	}

	server := handlers.NewJobsHandler(db, ":3000")
	server.Run()
}
