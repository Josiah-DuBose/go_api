package models

// User type structure
type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// Name returns User's username
func (u User) Name() string {
	return u.Username
}
