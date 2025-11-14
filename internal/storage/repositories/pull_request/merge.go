package repo_pull_request

import (
	"context"

	"github.com/sariya23/manage_pr_service/internal/models/domain"
)

func (r *PullRequestRepository) MergePullRequest(ctx context.Context, prID string) (*domain.PullRequest, error) {
	const operation = "storage.repositories.pull_request.MergePullRequest"

	//	updatePullRequestSQL := `update pull_request set status='MERGED' where pull_request_id = $1
	//returning pull_request_id, pull_request_name, author_id, status, merged_at, created_at`

	panic("implement me")
}
