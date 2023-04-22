package types

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type RegisterUser struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type LoginUser struct {
	Email      string `json:"email"`
	Password string `json:"password"`
}
