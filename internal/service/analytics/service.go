package serviceanalytics

import (
	"context"
	"log/slog"

	"github.com/sariya23/manage_pr_service/internal/models/dto"
)

type PullRequestRepository interface {
	GroupPullRequestsByAssignedReviewer(ctx context.Context) ([]dto.PullRequestAssignedReviewer, error)
}

type AnalyticsService struct {
	log             *slog.Logger
	pullRequestRepo PullRequestRepository
}

func NewAnalyticsService(log *slog.Logger, pullRequestRepo PullRequestRepository) *AnalyticsService {
	return &AnalyticsService{log: log, pullRequestRepo: pullRequestRepo}
}
