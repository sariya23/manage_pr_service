package repo_pull_request

import (
	"log/slog"

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
	log  *slog.Logger
}

func NewPullRequestRepository(conn *database.Database, log *slog.Logger) *PullRequestRepository {
	return &PullRequestRepository{
		conn: conn,
		log:  log,
	}
}
