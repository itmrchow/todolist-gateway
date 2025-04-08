package dto

import (
	"time"
)

type LoginUserRespDTO struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Token     string    `json:"token"`
	ExpiresIn time.Time `json:"expires_in"`
}
