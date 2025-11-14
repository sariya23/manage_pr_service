package repo_pull_request

import (
	"context"
	"fmt"

	"github.com/sariya23/manage_pr_service/internal/converters"
	"github.com/sariya23/manage_pr_service/internal/models/domain"
	"github.com/sariya23/manage_pr_service/internal/models/dto"
)

func (r *PullRequestRepository) CreatePullRequestAndAssignReviewers(ctx context.Context, prData dto.CreatePullRequestDTO, reviewerIDs []string) (*domain.PullRequest, error) {
	const operationPlace = "storage.repositories.pull_request.create_pull_request_and_assign_reviewers"

	insertPullRequestSQL := `insert into pull_request 
    (pull_request_id, pull_request_name, author_id, assigned_reviewers) values ($1, $2, $3, $4) 
    returning pull_request_id, pull_request_name, author_id, status, created_at, assigned_reviewers`

	var pullRequestDb dto.PullRequestDB

	prRow := r.conn.GetPool().QueryRow(ctx, insertPullRequestSQL, prData.ID, prData.Name, prData.AuthorID, reviewerIDs)
	err := prRow.Scan(
		&pullRequestDb.ID,
		&pullRequestDb.Name,
		&pullRequestDb.AuthorID,
		&pullRequestDb.Status,
		&pullRequestDb.CreatedAt,
		&pullRequestDb.AssignedReviewerIDs,
	)
	if err != nil {
		return nil, fmt.Errorf("%s:insertPullRequestSQL:%w", operationPlace, err)
	}
	res := converters.PullRequestDBToDomain(pullRequestDb)
	return &res, nil
}
