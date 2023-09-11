package repo

import (
	"database/sql"
	"fmt"

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
	return err
}

func (us UserRepo) GetUserByCriterion(criterion, value string) (*entity.User, error) {
	query := fmt.Sprintf("SELECT * FROM users WHERE %s = $1", criterion)
	rows, err := us.db.Query(query, value)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		return scanUser(rows)
	}

	return nil, fmt.Errorf("user with %s '%s' not found", criterion, value)
}

func scanUser(rows *sql.Rows) (*entity.User, error) {
	user := new(entity.User)
	err := rows.Scan(&user.ID, &user.Name, &user.Password, &user.Email)
	return user, err
}
