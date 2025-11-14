package repo_pull_request

import (
	"context"
	"fmt"

	"github.com/sariya23/manage_pr_service/internal/converters"
	"github.com/sariya23/manage_pr_service/internal/models/domain"
	"github.com/sariya23/manage_pr_service/internal/models/dto"
)

func (r *PullRequestRepository) MergePullRequest(ctx context.Context, prID string) (*domain.PullRequest, error) {
	const operationPlace = "storage.repositories.pull_request.MergePullRequest"

	updatePullRequestSQL := `update pull_request set status='MERGED' where pull_request_id = $1
	returning pull_request_id, pull_request_name, author_id, status, merged_at, created_at, assigned_reviewers`

	row := r.conn.GetPool().QueryRow(ctx, updatePullRequestSQL, prID)
	var pullRequest dto.PullRequestDB

	err := row.Scan(
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

	res := converters.PullRequestDBToDomain(pullRequest)
	return &res, nil
}
