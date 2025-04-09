package service

import (
	"github.com/nathanSeixeiro/gateway-payments/internal/domain"
	"github.com/nathanSeixeiro/gateway-payments/internal/dto"
)

type AccountService struct {
	accountRepository domain.AccountRepository
}

func NewAccountService(repository domain.AccountRepository) *AccountService {
	return &AccountService{accountRepository: repository}
}

func (s *AccountService) CreateAccount(req dto.CreateAccountDTO) (*dto.AccountResponseDTO, error) {
	account := dto.ToDomain(&req)
	existingAccount, err := s.accountRepository.FindByAPIKey(account.APIKey)
	if err != nil && err != domain.ErrAccountNotFound {
		return nil, err
	}
	if existingAccount != nil {
		return nil, domain.ErrDuplicateAPIKey
	}
	err = s.accountRepository.Save(account); if err != nil {
		return nil, err
	}
	res := dto.FromDomain(account)
	return &res, nil
}

func (s *AccountService) UpdateBalance(apiKey string, amount float64) (*dto.AccountResponseDTO, error) {
	account, err := s.accountRepository.FindByAPIKey(apiKey)
	if err != nil {
		return nil, err
	}
	account.UpdateBalance(amount)
	err = s.accountRepository.UpdateBalance(account)
	if err != nil {
		return nil, err
	}
	res := dto.FromDomain(account)
	return &res, nil
}

func (s *AccountService) FindByAPIKey(apiKey string) (*dto.AccountResponseDTO, error) {
	account, err := s.accountRepository.FindByAPIKey(apiKey)
	if err != nil {
		return nil, err
	}
	res := dto.FromDomain(account)
	return &res, nil
}

func (s *AccountService) FindByID(id string) (*dto.AccountResponseDTO, error) {
	account, err := s.accountRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	res := dto.FromDomain(account)
	return &res, nil
}
