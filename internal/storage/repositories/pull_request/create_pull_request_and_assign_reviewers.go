package repo_pull_request

import (
	"context"
	"fmt"
	"strings"

	"github.com/sariya23/manage_pr_service/internal/models/domain"
	"github.com/sariya23/manage_pr_service/internal/models/dto"
)

func (r *PullRequestRepository) CreatePullRequestAndAssignReviewers(ctx context.Context, dto dto.CreatePullRequestDTO, reviewerIDs []string) (*domain.PullRequest, error) {
	const operationPlace = "storage.repositories.pull_request.create_pull_request_and_assign_reviewers"

	insertPullRequestSQL := `insert into pull_request 
    (pull_request_id, pull_request_name, author_id) values ($1, $2, $3) 
    returning pull_request_id, pull_request_name, author_id, status, created_at`

	insertReviewersSQL := strings.Builder{}
	insertReviewersSQL.WriteString(`insert into reviewer_pull_request (user_id, pull_request_id) values `)
	insertReviewersArgs := make([]any, 0, len(reviewerIDs)*2)
	insertReviewersValues := make([]string, 0, len(reviewerIDs))

	for _, reviewerID := range reviewerIDs {
		insertReviewersValues = append(insertReviewersValues, fmt.Sprintf("($%d, $%d)",
			len(insertReviewersArgs),
			len(insertReviewersArgs)+1))
		insertReviewersArgs = append(insertReviewersArgs, reviewerID, dto.ID)
	}
	insertReviewersSQL.WriteString(strings.Join(insertReviewersValues, ", "))

	var pullRequest domain.PullRequest
	prRow := r.conn.GetPool().QueryRow(ctx, insertPullRequestSQL, dto.ID, dto.Name, dto.AuthorID)
	err := prRow.Scan(
		&pullRequest.ID,
		&pullRequest.Name,
		&pullRequest.AuthorID,
		&pullRequest.Status,
		&pullRequest.CreatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("%s:insertPullRequestSQL:%w", operationPlace, err)
	}
	if len(reviewerIDs) > 0 {
		_, err = r.conn.GetPool().Exec(ctx, insertReviewersSQL.String(), insertReviewersArgs...)
		if err != nil {
			return nil, fmt.Errorf("%s:insertReviewersSQL:%w", operationPlace, err)
		}
	}
	return &pullRequest, nil
}
