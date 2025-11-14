package service_pull_request

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/sariya23/manage_pr_service/internal/models/domain"
	"github.com/sariya23/manage_pr_service/internal/outerror"
)

func (s *PullRequestService) Merge(ctx context.Context, prID string) (*domain.PullRequest, error) {
	const operationPlace = "service.pull_request.Merge"
	log := s.log.With("operationPlace", operationPlace)

	pr, err := s.PullRequestRepo.GetPullRequest(ctx, prID)
	if err != nil {
		if errors.Is(err, outerror.ErrPullRequestNotFound) {
			log.Warn("pull request not found", slog.String("pr_id", prID))
			return nil, fmt.Errorf("%s:%w", operationPlace, err)
		}
		log.Error("failed to get pull request", slog.String("pr_id", prID), slog.String("error", err.Error()))
		return nil, fmt.Errorf("%s:%w", operationPlace, err)
	}

	if pr.MergedAt != nil {
		return pr, nil
	}

	pr, err = s.PullRequestRepo.MergePullRequest(ctx, prID)
	if err != nil {
		if errors.Is(err, outerror.ErrPullRequestNotFound) {
			log.Warn("pull request not found", slog.String("pr_id", prID))
			return nil, fmt.Errorf("%s:%w", operationPlace, err)
		}
		log.Error("failed to merge pull request", slog.String("pr_id", prID), slog.String("error", err.Error()))
		return nil, fmt.Errorf("%s:%w", operationPlace, err)
	}
	return pr, nil
}
