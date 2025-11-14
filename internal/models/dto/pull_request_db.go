package dto

import (
	"database/sql"
)

type PullRequestDB struct {
	ID                  string
	Name                string
	AuthorID            string
	Status              string
	MergedAt            sql.NullTime
	CreatedAt           sql.NullTime
	AssignedReviewerIDs []string
}
