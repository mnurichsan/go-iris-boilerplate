package model

type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserResponse struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GetUserResponse struct {
	Id    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
