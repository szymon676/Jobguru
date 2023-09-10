package repo

import (
	"testing"
	"time"

	"github.com/szymon676/betterdocker/postgres"
	"github.com/szymon676/jobguru/internal/entity"
	"github.com/szymon676/jobguru/internal/migrate"
)

func TestRepo(t *testing.T) {
	opts := &postgres.PostgresContainerOptions{}
	container := postgres.NewPostgresContainer(opts)
	container.Run()

	dsn := "host=localhost user=postgres password=password dbname=postgres port=5432 sslmode=disable"
	db := migrate.MigratePostgresDB(dsn)

	jr := NewJobRepo(db)

	job := &entity.Job{
		UserID:      1,
		Title:       "Software Engineer",
		Company:     "Example Inc",
		Skills:      []string{"Go", "SQL"},
		Salary:      100000,
		Description: "A software engineering job",
		Currency:    "USD",
		Date:        time.Now(),
		Location:    "New York",
	}

	t.Run("create job test", func(t *testing.T) {
		err := jr.CreateJob(job)
		if err != nil {
			t.Fatal("shall not return error.")
		}
	})
	t.Run("get all jobs test", func(t *testing.T) {
		jobs, err := jr.GetJobs()

		if err != nil {
			t.Fatalf("Expected no error, but got %v", err)
		}

		if len(jobs) != 2 {
			t.Fatalf("Expected 2 jobs, but got %d", len(jobs))
		}
	})
	t.Run("update job test", func(t *testing.T) {
		//
	})

	container.Stop()
}
