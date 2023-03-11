package main

import (
	"flag"
	"fmt"

	"github.com/szymon676/job-guru/jobs/internal/database"
	"github.com/szymon676/job-guru/jobs/internal/handlers"
)

func main() {
	dsn := flag.String("dsn", "host=localhost user=postgres password=1234 dbname=jobguru port=5432 sslmode=disable", "dsn to connect to the database")
	db, err := database.NewDatabase("postgres", *dsn)

	if err != nil {
		fmt.Printf("Error creating database")
	}

	server := handlers.NewJobsHandler(db, ":4000")
	server.Run()
}
