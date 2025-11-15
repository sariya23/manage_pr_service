package pullrequest

import (
	"github.com/sariya23/manage_pr_service/internal/storage/database"
)

type PullRequestRepository struct {
	conn *database.Database
}

func NewPullRequestRepository(conn *database.Database) *PullRequestRepository {
	return &PullRequestRepository{
		conn: conn,
	}
}
