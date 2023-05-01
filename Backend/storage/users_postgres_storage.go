package storage

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/szymon676/jobguru/types"
)

type PostgreUserStorage struct {
	db *sql.DB
}

func NewPostgreUserStorage(db *sql.DB) *PostgreUserStorage {
	return &PostgreUserStorage{
		db: db,
	}
}

func (us PostgreUserStorage) CreateUser(name, password, email string) error {
	query := "INSERT INTO users (username, password, email) VALUES($1, $2, $3)"

	_, err := us.db.Exec(query, name, password, email)
	if err != nil {
		return err
	}

	return nil
}

func (us PostgreUserStorage) GetUserByID(id int) (*types.User, error) {
	rows, err := us.db.Query("select * from users where id = $1", id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return scanUser(rows)
	}

	return nil, fmt.Errorf("user %d not found", id)
}

func (us PostgreUserStorage) GetUserByEmail(email string) (*types.User, error) {
	rows, err := us.db.Query("select * from users where email = $1", email)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return scanUser(rows)
	}

	return nil, fmt.Errorf("user %s not found", email)
}

func scanUser(rows *sql.Rows) (*types.User, error) {
	account := new(types.User)
	err := rows.Scan(
		&account.ID,
		&account.Name,
		&account.Password,
		&account.Email,
	)

	return account, err
}
