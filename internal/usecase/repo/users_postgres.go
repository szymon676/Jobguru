package repo

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/szymon676/jobguru/internal/entity"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (us UserRepo) CreateUser(req entity.RegisterUser) error {
	query := "INSERT INTO users (username, password, email) VALUES($1, $2, $3)"

	_, err := us.db.Exec(query, req.Name, req.Password, req.Email)
	if err != nil {
		return err
	}

	return nil
}

func (us UserRepo) GetUserByID(id int) (*entity.User, error) {
	rows, err := us.db.Query("select * from users where id = $1", id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return scanUser(rows)
	}

	return nil, fmt.Errorf("user %d not found", id)
}

func (us UserRepo) GetUserByEmail(email string) (*entity.User, error) {
	rows, err := us.db.Query("select * from users where email = $1", email)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return scanUser(rows)
	}

	return nil, fmt.Errorf("user %s not found", email)
}

func scanUser(rows *sql.Rows) (*entity.User, error) {
	account := new(entity.User)
	err := rows.Scan(
		&account.ID,
		&account.Name,
		&account.Password,
		&account.Email,
	)

	return account, err
}
