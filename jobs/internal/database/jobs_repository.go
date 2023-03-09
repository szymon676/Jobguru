package database

import (
	"fmt"

	"github.com/lib/pq"
)

func InsertUser(title, category, description string, skills []string) error {
	query := "INSERT INTO jobs (title, category, skills, description) VALUES($1, $2, $3, $4)"
	convskills := pq.Array(skills)
	_, err := DB.Query(query, title, category, convskills, description)
	if err != nil {
		return fmt.Errorf("insert into jobs failed: %v", err)
	}

	return nil
}
