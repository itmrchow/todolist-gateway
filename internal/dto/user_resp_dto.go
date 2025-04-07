package dto

import (
	"time"

	"github.com/google/uuid"
)

type RegisterUserRespDTO struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Token     string    `json:"token"`
	ExpiresIn time.Time `json:"expires_in"`
}
