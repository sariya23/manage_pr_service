package apiusers

import (
	"context"
	"log/slog"

	"github.com/sariya23/manage_pr_service/internal/models/domain"
)

type userService interface {
	SetIsActive(ctx context.Context, userID string, isActive bool) (*domain.User, error)
	GetReviews(ctx context.Context, userID string) ([]domain.PullRequest, error)
	GetUserTeam(ctx context.Context, userID string) (string, error)
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
