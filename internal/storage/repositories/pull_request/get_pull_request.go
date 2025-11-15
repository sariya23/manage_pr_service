package pullrequest

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/sariya23/manage_pr_service/internal/converters"
	"github.com/sariya23/manage_pr_service/internal/models/domain"
	"github.com/sariya23/manage_pr_service/internal/models/dto"
	"github.com/sariya23/manage_pr_service/internal/outerror"
)

func (r *PullRequestRepository) GetPullRequest(ctx context.Context, prID string) (*domain.PullRequest, error) {
	const operationPlace = "storage.repositories.pull_request.get_pull_request"

	getPullRequestSQL := `select pull_request_id, pull_request_name, 
author_id, status, merged_at, created_at, assigned_reviewers from pull_request where pull_request_id=$1`

	var pullRequestDB dto.PullRequestDB

	row := r.conn.GetPool().QueryRow(ctx, getPullRequestSQL, prID)
	err := row.Scan(
		&pullRequestDB.ID,
		&pullRequestDB.Name,
		&pullRequestDB.AuthorID,
		&pullRequestDB.Status,
		&pullRequestDB.MergedAt,
		&pullRequestDB.CreatedAt,
		&pullRequestDB.AssignedReviewerIDs)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%s:%w", operationPlace, outerror.ErrPullRequestNotFound)
		}
		return nil, fmt.Errorf("%s: %w", operationPlace, err)
	}

	res := converters.PullRequestDBToDomain(pullRequestDB)
	return &res, nil
}
