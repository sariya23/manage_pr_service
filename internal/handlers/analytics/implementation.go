package analytics

import (
	"context"
	"log/slog"
)

type PullRequestService interface {
	GroupPullRequestsByAssignedReviewer(ctx context.Context) (map[string]int, error)
}

type AnalyticsImplementation struct {
	log                *slog.Logger
	PullRequestService PullRequestService
}

func NewAnalyticsImplementation(log *slog.Logger, pullRequestRepo PullRequestService) *AnalyticsImplementation {
	return &AnalyticsImplementation{
		log:                log,
		PullRequestService: pullRequestRepo,
	}
}
