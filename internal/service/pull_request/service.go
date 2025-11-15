package pullrequest

import (
	"context"
	"log/slog"

	"github.com/sariya23/manage_pr_service/internal/models/domain"
	"github.com/sariya23/manage_pr_service/internal/models/dto"
)

type UserRepository interface {
	GetUserByID(ctx context.Context, userID string) (*domain.User, error)
}

type PullRequestRepository interface {
	GetPullRequest(ctx context.Context, prID string) (*domain.PullRequest, error)
	CreatePullRequestAndAssignReviewers(ctx context.Context, dtoPR dto.CreatePullRequestDTO, reviewerIDs []string) (*domain.PullRequest, error)
	MergePullRequest(ctx context.Context, prID string) (*domain.PullRequest, error)
	ReassignPullRequest(ctx context.Context, prID string, oldReviewerID, newReviewerID string) (*domain.PullRequest, error)
}

type TeamRepository interface {
	GetUserTeam(ctx context.Context, userID string) (string, error)
	GetTeamMembers(ctx context.Context, teamName string) ([]domain.User, error)
}

type PullRequestService struct {
	log             *slog.Logger
	PullRequestRepo PullRequestRepository
	UserRepo        UserRepository
	TeamRepo        TeamRepository
}

func NewPullRequestService(log *slog.Logger, prRepo PullRequestRepository, userRepo UserRepository, teamRepo TeamRepository) *PullRequestService {
	return &PullRequestService{
		log:             log,
		PullRequestRepo: prRepo,
		UserRepo:        userRepo,
		TeamRepo:        teamRepo,
	}
}
