package serviceanalytics

import (
	"context"
	"fmt"

	"github.com/sariya23/manage_pr_service/internal/models/dto"
)

func (s *AnalyticsService) GroupPullRequestsByAssignedReviewer(ctx context.Context) (map[string]int, error) {
	const operationPlace = "service.analytics.GroupPullRequestsByAssignedReviewer"

	groups, err := s.pullRequestRepo.GroupPullRequestsByAssignedReviewer(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", operationPlace, err)
	}
	m := dto.ToMapPullRequestAssignedReviewer(groups)
	return m, nil
}
