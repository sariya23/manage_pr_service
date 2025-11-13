package apiusers

import (
	"context"
	"log/slog"

	"github.com/sariya23/manage_pr_service/internal/models/domain"
)

type userService interface {
	SetIsActive(ctx context.Context, userId string, isActive bool) (*domain.User, error)
	GetReviews(ctx context.Context, userId string) ([]domain.PullRequest, error)
}

type UsersImplementation struct {
	logger      *slog.Logger
	userService userService
}

func NewUsersImplementation(logger *slog.Logger, userService userService) *UsersImplementation {
	return &UsersImplementation{
		logger:      logger,
		userService: userService,
	}
}
