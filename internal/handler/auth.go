package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/mail"

	"github.com/nseve/first-go-restapi/internal/response"
	"github.com/nseve/first-go-restapi/internal/service"
)

type AuthHandler struct {
	service *service.AuthService
}

func NewAuthHandler(service *service.AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type MessageResponse struct {
	Message string `json:"message"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func (h *AuthHandler) Register(
	w http.ResponseWriter,
	r *http.Request,
) {
	var req AuthRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteError(
			w,
			http.StatusBadRequest,
			"invalid request body",
		)
		return
	}

	_, err := mail.ParseAddress(req.Email)
	if err != nil || len(req.Password) < 6 {
		response.WriteError(
			w,
			http.StatusBadRequest,
			"invalid email or password",
		)
		return
	}

	err = h.service.Register(
		req.Email,
		req.Password,
	)

	if errors.Is(err, service.ErrUserAlreadyExists) {
		response.WriteError(
			w,
			http.StatusConflict,
			"user already exists",
		)
		return
	}

	if err != nil {
		response.WriteError(
			w,
			http.StatusInternalServerError,
			"failed to create user",
		)
		return
	}

	response.WriteJSON(
		w,
		http.StatusCreated,
		MessageResponse{
			Message: "user created",
		},
	)
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req AuthRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	token, err := h.service.Login(req.Email, req.Password)

	if errors.Is(err, service.ErrInvalidCredentials) {
		response.WriteError(
			w,
			http.StatusUnauthorized,
			"invalid credentials",
		)
		return
	}

	if err != nil {
		response.WriteError(
			w,
			http.StatusInternalServerError,
			"login failed",
		)
		return
	}

	response.WriteJSON(
		w,
		http.StatusOK,
		LoginResponse{
			Token: token,
		},
	)
}
