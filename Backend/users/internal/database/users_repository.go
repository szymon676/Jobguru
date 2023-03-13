package database

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
