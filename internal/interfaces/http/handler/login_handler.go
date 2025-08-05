package handler

import (
	"context"
	"encoding/json"
	"github.com/rfanazhari/factory-login/internal/application/dto"
	"github.com/rfanazhari/factory-login/internal/application/usecase"
	"net/http"
)

// LoginHandler handles HTTP requests for login
type LoginHandler struct {
	loginUseCase *usecase.LoginUseCase
}

func NewLoginHandler(loginUseCase *usecase.LoginUseCase) *LoginHandler {
	return &LoginHandler{loginUseCase: loginUseCase}
}

func (h *LoginHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req dto.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	response, err := h.loginUseCase.Execute(context.Background(), &req)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if response.Success {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}

	json.NewEncoder(w).Encode(response)
}
