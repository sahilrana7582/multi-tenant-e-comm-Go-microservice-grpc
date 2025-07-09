package service

import (
	"context"
	"net/http"

	"github.com/sahilrana7582/multi-tenent-e-com-user-service/internal/domain"
	"github.com/sahilrana7582/multi-tenent-e-com-user-service/internal/dto"
	"github.com/sahilrana7582/multi-tenent-e-com-user-service/internal/repository"
	"github.com/sahilrana7582/multi-tenent-e-com-user-service/internal/utils"
)

type UserService interface {
	RegisterUser(ctx context.Context, tenantID string, req dto.RegisterUserRequest) (*domain.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) RegisterUser(ctx context.Context, tenantID string, req dto.RegisterUserRequest) (*domain.User, error) {

	hashedPwd, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, utils.StatusError{
			Er:     err,
			Status: http.StatusInternalServerError,
		}
	}

	user := &domain.User{
		TenantID: tenantID,
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPwd,
		Role:     "user",
	}

	if err := s.repo.Create(ctx, user); err != nil {
		return nil, utils.StatusError{
			Er:     err,
			Status: http.StatusInternalServerError,
		}
	}

	return user, nil
}
