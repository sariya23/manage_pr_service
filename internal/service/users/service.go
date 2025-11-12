package serviceusers

import (
	"context"
	"log/slog"

	"github.com/sariya23/manage_pr_service/internal/models"
)

type UserRepository interface {
	SetIsActive(ctx context.Context, id int64, isActive bool) (*models.User, error)
	GetUserByID(ctx context.Context, userID int64) (*models.User, error)
}

type ReviewUserPullRequestRepository interface {
	GetUserReviews(ctx context.Context, userID int64) ([]models.PullRequest, error)
}

type UsersService struct {
	log            *slog.Logger
	userRepo       UserRepository
	reviewUserRepo ReviewUserPullRequestRepository
}

func NewUsersService(log *slog.Logger, userRepo UserRepository, reviewUserRepo ReviewUserPullRequestRepository) *UsersService {
	return &UsersService{
		log:            log,
		userRepo:       userRepo,
		reviewUserRepo: reviewUserRepo,
	}
}
