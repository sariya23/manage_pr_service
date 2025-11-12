package serviceusers

import (
	"context"
	"log/slog"

	"github.com/sariya23/manage_pr_service/internal/models"
)

type UserRepository interface {
	SetIsActive(ctx context.Context, id int64, isActive bool) (*models.User, error)
}

type UsersService struct {
	log      *slog.Logger
	userRepo UserRepository
}

func NewUsersService(log *slog.Logger, userRepo UserRepository) *UsersService {
	return &UsersService{
		log:      log,
		userRepo: userRepo,
	}
}
