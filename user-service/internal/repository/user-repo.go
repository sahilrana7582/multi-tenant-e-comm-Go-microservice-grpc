package repository

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sahilrana7582/multi-tenent-e-com-user-service/internal/domain"
)

type UserRepository interface {
	Create(ctx context.Context, user *domain.User) error
}

type userRepo struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) UserRepository {
	return &userRepo{db: db}
}

func (r *userRepo) Create(ctx context.Context, user *domain.User) error {
	query := `
		INSERT INTO users (tenant_id, name, email, password, role, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id
	`

	err := r.db.QueryRow(ctx, query,
		user.TenantID,
		user.Name,
		user.Email,
		user.Password,
		user.Role,
		time.Now(),
		time.Now(),
	).Scan(&user.ID)

	return err
}
