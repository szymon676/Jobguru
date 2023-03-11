package main

import (
	"flag"
	"fmt"

	"github.com/szymon676/job-guru/jobs/internal/database"
	"github.com/szymon676/job-guru/jobs/internal/handlers"
)

func main() {
	dsn := flag.String("dsn", "host=localhost user=postgres password=1234 dbname=jobguru port=5432 sslmode=disable", "dsn to connect to the database")
	if _, err := database.NewDatabase("postgres", *dsn); err != nil {
		fmt.Printf("Error creating database: %v", err)
	}

	server := handlers.NewJobsHandler(":1337")
	server.Run()
}
