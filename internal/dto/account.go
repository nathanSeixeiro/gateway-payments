package dto

import (
	"time"

	"github.com/nathanSeixeiro/gateway-payments/internal/domain"
)

type CreateAccountDTO struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type AccountResponseDTO struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Balance   float64   `json:"balance"`
	APIKey    string    `json:"api_key,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ToDomain(dto *CreateAccountDTO) *domain.Account {
	return domain.NewAccount(dto.Name, dto.Email)
}

func FromDomain(acc *domain.Account) AccountResponseDTO {
	return AccountResponseDTO{
		ID:        acc.ID,
		Name:      acc.Name,
		Email:     acc.Email,
		Balance:   acc.Balance,
		APIKey:    acc.APIKey,
		CreatedAt: acc.CreatedAt,
		UpdatedAt: acc.UpdatedAt,
	}
}
