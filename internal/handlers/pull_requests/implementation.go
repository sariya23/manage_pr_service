package api_pull_requests

import (
	"context"
	"log/slog"

	"github.com/sariya23/manage_pr_service/internal/models/domain"
	"github.com/sariya23/manage_pr_service/internal/models/dto"
)

type PullRequestService interface {
	CreatePullRequestAndAssignReviewers(ctx context.Context, prData dto.CreatePullRequestDTO) (*domain.PullRequest, error)
	Merge(ctx context.Context, pdID string) (*domain.PullRequest, error)
	Reassign(ctx context.Context, prID string, oldReviewerID string) (*domain.PullRequest, string, error)
}

type PullRequestImplementation struct {
	logger    *slog.Logger
	prService PullRequestService
}

func NewPullRequestImplementation(log *slog.Logger, prService PullRequestService) *PullRequestImplementation {
	return &PullRequestImplementation{
		logger:    log,
		prService: prService,
	}
}
