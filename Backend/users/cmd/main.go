package main

import (
	"flag"

	_ "github.com/lib/pq"
	database "github.com/szymon676/job-guru/users/domain/repository"
	api "github.com/szymon676/job-guru/users/domain/transport/http"
)

func main() {
	dsn := flag.String("dsn", "host=localhost user=postgres password=1234 dbname=jobguru-users port=5432 sslmode=disable", "dsn to connect to the database")
	flag.Parse()

	if _, err := database.NewDatabase("postgres", *dsn); err != nil {
		panic(err)
	}

	server := api.NewApiServer(":5000")
	server.Run()
}
