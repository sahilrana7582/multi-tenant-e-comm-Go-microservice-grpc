package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/sahilrana7582/multi-tenent-e-com-user-service/internal/dto"
	"github.com/sahilrana7582/multi-tenent-e-com-user-service/internal/service"
	"github.com/sahilrana7582/multi-tenent-e-com-user-service/internal/utils"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) error {
	tenantID := r.Header.Get("tenend_id")
	if tenantID == "" {
		return utils.NewStatusError(
			errors.New("missing error id"),
			http.StatusBadRequest,
		)
	}

	var req dto.RegisterUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return utils.NewStatusError(
			err,
			http.StatusBadRequest,
		)
	}

	user, err := h.userService.RegisterUser(r.Context(), tenantID, req)
	if err != nil {
		return err
	}

	return utils.WriteJSON(w, http.StatusCreated, "User registered successfully", user)
}
