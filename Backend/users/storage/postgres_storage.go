package storage

import (
	"database/sql"
	"fmt"

	"github.com/szymon676/job-guru/users/types"
)

var DB *sql.DB

type PostgreStorage struct {
	db *sql.DB
}

func NewDatabase(driverName, dsn string) (*sql.DB, error) {
	db, err := sql.Open(driverName, dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS users (
            id SERIAL PRIMARY KEY,
            username text,
			password text,
			email text
		);
    `)

	DB = db

	return nil, nil
}

func NewPostgresStorage() (*PostgreStorage, error) {
	return &PostgreStorage{
		db: DB,
	}, nil
}

func (ps PostgreStorage) CreateUser(name, password, email string) error {
	query := "INSERT INTO users (username, password, email) VALUES($1, $2, $3)"

	_, err := ps.db.Exec(query, name, password, email)
	if err != nil {
		return err
	}

	return nil
}

func (ps PostgreStorage) GetUserByID(id int) (*types.User, error) {
	rows, err := ps.db.Query("select * from users where id = $1", id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return scanUser(rows)
	}

	return nil, fmt.Errorf("user %d not found", id)
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
