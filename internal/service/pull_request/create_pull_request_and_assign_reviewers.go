package service_pull_request

import (
	"context"

	"github.com/sariya23/manage_pr_service/internal/models/domain"
	"github.com/sariya23/manage_pr_service/internal/models/dto"
)

func (s *PullRequestService) CreatePullRequestAndAssignReviewers(ctx context.Context, prData dto.CreatePullRequestDTO) (domain.PullRequest, []domain.User, error) {
	
}
