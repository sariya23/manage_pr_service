package serviceusers

import (
	"context"
	"log/slog"

	"github.com/sariya23/manage_pr_service/internal/models/domain"
)

type UserRepository interface {
	SetIsActive(ctx context.Context, userID string, isActive bool) (*domain.User, error)
	GetUserByID(ctx context.Context, userID string) (*domain.User, error)
}

type ReviewUserPullRequestRepository interface {
	GetUserReviews(ctx context.Context, userID string) ([]domain.PullRequest, error)
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
