package pullrequest

import (
	"context"
	"fmt"

	"github.com/sariya23/manage_pr_service/internal/models/dto"
)

func (r *PullRequestRepository) GroupPullRequestsByAssignedReviewer(ctx context.Context) ([]dto.PullRequestAssignedReviewer, error) {
	const operationPlace = "storage.repositories.pull_request.groupPullRequestsByAssignedReviewer"

	query := `select u.user_id as user_id, coalesce(r.reviews, 0) as reviews
from "user" u left join
(select unnest(assigned_reviewers) as user_id, count(*) as reviews from pull_request group by user_id) r
on r.user_id = u.user_id;`

	rows, err := r.conn.GetPool().Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", operationPlace, err)
	}
	defer rows.Close()
	var groups []dto.PullRequestAssignedReviewer
	for rows.Next() {
		var row dto.PullRequestAssignedReviewer
		err = rows.Scan(
			&row.UserID,
			&row.CountPullRequestID)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", operationPlace, err)
		}
		if rows.Err() != nil {
			return nil, fmt.Errorf("%s: %w", operationPlace, rows.Err())
		}
		groups = append(groups, row)
	}
	return groups, nil
}
