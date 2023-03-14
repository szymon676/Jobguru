package main

import (
	"flag"

	_ "github.com/lib/pq"
	"github.com/szymon676/job-guru/jobs/domain/repository"
	api "github.com/szymon676/job-guru/jobs/domain/transport/http"
)

func main() {
	dsn := flag.String("dsn", "host=localhost user=postgres password=1234 dbname=jobguru-jobs port=5432 sslmode=disable", "dsn to connect to the database")
	flag.Parse()

	if _, err := repository.NewDatabase("postgres", *dsn); err != nil {
		panic(err)
	}

	server := api.NewApiServer(":5001")
	server.Run()
}
