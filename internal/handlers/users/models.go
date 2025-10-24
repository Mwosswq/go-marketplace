package users

type UserResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type UserCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
