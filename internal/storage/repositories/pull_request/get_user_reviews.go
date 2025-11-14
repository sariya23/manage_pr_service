package repo_pull_request

import (
	"context"
	"fmt"

	"github.com/sariya23/manage_pr_service/internal/converters"
	"github.com/sariya23/manage_pr_service/internal/models/domain"
	"github.com/sariya23/manage_pr_service/internal/models/dto"
)

func (r *PullRequestRepository) GetUserReviews(ctx context.Context, userID string) ([]domain.PullRequest, error) {
	const operationPlace = "storage.repositories.pull_request.GetUserReviews"

	getPullRequestsSQL := `select pull_request_id, pull_request_name, 
author_id, status, created_at, merged_at, assigned_reviewers from pull_request where $1=any(assigned_reviewers)`
	prRows, err := r.conn.GetPool().Query(ctx, getPullRequestsSQL, userID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", operationPlace, err)
	}
	defer prRows.Close()
	var pullRequests []dto.PullRequestDB
	for prRows.Next() {
		var pullRequest dto.PullRequestDB
		err = prRows.Scan(
			&pullRequest.ID,
			&pullRequest.Name,
			&pullRequest.AuthorID,
			&pullRequest.Status,
			&pullRequest.MergedAt,
			&pullRequest.CreatedAt,
			&pullRequest.AssignedReviewerIDs)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", operationPlace, err)
		}
		if prRows.Err() != nil {
			return nil, fmt.Errorf("%s: %w", operationPlace, prRows.Err())
		}
		pullRequests = append(pullRequests, pullRequest)
	}

	pullRequestsRes := converters.MultiPullRequestDBToDomain(pullRequests)
	return pullRequestsRes, nil
}
