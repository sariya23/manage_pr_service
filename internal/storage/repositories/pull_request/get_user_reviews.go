package repo_pull_request

import (
	"context"
	"fmt"

	"github.com/sariya23/manage_pr_service/internal/models/domain"
)

func (r *PullRequestRepository) GetUserReviews(ctx context.Context, userID string) ([]domain.PullRequest, error) {
	const operationPlace = "storage.repositories.pull_request.GetUserReviews"

	getPullRequestsSQL := fmt.Sprintf("select %s, %s, %s, %s, %s, %s from %s join %s using(%s) where %s=$1",
		PullRequestTablePullRequestIDField,
		PullRequestRTablePullRequestNameField,
		PullRequestTableAuthorIDField,
		PullRequestTableStatusField,
		PullRequestTableMergedField,
		PullRequestTableCreatedField,
		PullRequestTableName,
		UserPullRequestTableName,
		PullRequestTablePullRequestIDField,
		UserPullRequestTableUserID)
	prRows, err := r.conn.GetPool().Query(ctx, getPullRequestsSQL, userID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", operationPlace, err)
	}
	defer prRows.Close()
	var pullRequests []domain.PullRequest
	for prRows.Next() {
		var pullRequest domain.PullRequest
		err = prRows.Scan(
			&pullRequest.ID,
			&pullRequest.Name,
			&pullRequest.AuthorID,
			&pullRequest.Status,
			&pullRequest.MergedAt,
			&pullRequest.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", operationPlace, err)
		}
		if prRows.Err() != nil {
			return nil, fmt.Errorf("%s: %w", operationPlace, prRows.Err())
		}
		pullRequests = append(pullRequests, pullRequest)
	}
	return pullRequests, nil
}
