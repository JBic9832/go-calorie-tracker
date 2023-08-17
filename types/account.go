package types

type Account struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateAccountRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewAccount(username string, email string, password string) *Account {
	return &Account{
		Username: username,
		Email:    email,
		Password: password,
	}
}
