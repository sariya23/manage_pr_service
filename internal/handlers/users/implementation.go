package apiusers

import (
	"context"
	"log/slog"

	"github.com/sariya23/manage_pr_service/internal/models"
)

type userService interface {
	SetIsActive(ctx context.Context, userId int64, isActive bool) (models.User, error)
	GetUserReviews(ctx context.Context, userId int64) ([]models.PullRequest, error)
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
