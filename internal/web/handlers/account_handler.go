package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/nathanSeixeiro/gateway-payments/internal/dto"
	"github.com/nathanSeixeiro/gateway-payments/internal/service"
)

type AccountHandler struct {
	accountService *service.AccountService
}

func NewAccountHandler(accountService *service.AccountService) *AccountHandler {
	return &AccountHandler{accountService: accountService}
}

func (h *AccountHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input dto.CreateAccountDTO
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	output, err := h.accountService.CreateAccount(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
}

func (h *AccountHandler) Get(w http.ResponseWriter, r *http.Request) {
	apiKey := r.Header.Get("X-API-Key")
	if apiKey == "" {
		http.Error(w, "API key is required", http.StatusUnauthorized)
		return
	}

	output, err := h.accountService.FindByAPIKey(apiKey)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}
