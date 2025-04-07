package handlers

type RegisterUserReqDTO struct {
	Email    string `validate:"required,email" json:"email"`
	Name     string `validate:"required,min=1,max=20" json:"name"`
	Password string `validate:"required,min=6,max=20" json:"password"`
}

type LoginUserReqDTO struct {
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required,min=6,max=20" json:"password"`
}
