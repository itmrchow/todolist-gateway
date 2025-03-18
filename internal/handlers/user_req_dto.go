package handlers

type RegisterUserReqDTO struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type LoginUserReqDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
