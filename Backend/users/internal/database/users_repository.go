package database

import (
	"database/sql"
	"fmt"

	"github.com/szymon676/job-guru/users/internal/models"
)

//type User struct {
//	ID       int      `json:"id"`
//	Name     string   `json:"name"`
//	Password string   `json:"password"`
//	Email    string   `json:"email"`
//	Jobs     []string `json:"jobs"`
//}

func CreateUser(name, password, email string) error {
	query := "INSERT INTO users (username, password, email) VALUES($1, $2, $3)"

	_, err := DB.Exec(query, name, password, email)
	if err != nil {
		return err
	}

	return nil
}

func GetUserByID(id int) (*models.User, error) {
	rows, err := DB.Query("select * from users where id = $1", id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return scanUser(rows)
	}

	return nil, fmt.Errorf("user %d not found", id)
}

func scanUser(rows *sql.Rows) (*models.User, error) {
	account := new(models.User)
	err := rows.Scan(
		&account.ID,
		&account.Name,
		&account.Password,
		&account.Email,
	)
	return account, err
}
