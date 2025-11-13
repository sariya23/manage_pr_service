package repo_pull_request

import (
	"github.com/sariya23/manage_pr_service/internal/storage/database"
)

const (
	PullRequestTableName                  = "pull_request"
	PullRequestTablePullRequestIDField    = "pull_request_id"
	PullRequestRTablePullRequestNameField = "pull_request_name"
	PullRequestTableAuthorIDField         = "author_id"
	PullRequestTableStatusField           = "status"

	UserPullRequestTableName               = "user_pull_request"
	UserPullRequestTableUserID             = "user_id"
	UserPullRequestTablePullRequestIDField = "pull_request_id"
)

type PullRequestRepository struct {
	conn *database.Database
}

func NewPullRequestRepository(conn *database.Database) *PullRequestRepository {
	return &PullRequestRepository{
		conn: conn,
	}
}
