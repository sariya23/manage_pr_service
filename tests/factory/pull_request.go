//go:build integrations

package factory

import (
	"database/sql"
	"time"
)

type PullRequest struct {
	ID                  string
	Name                string
	AuthorID            string
	Status              string
	MergedAt            *time.Time
	CreatedAt           time.Time
	AssignedReviewerIDs []string
}

type PullRequestDB struct {
	ID                  string
	Name                string
	AuthorID            string
	Status              string
	MergedAt            sql.NullTime
	CreatedAt           time.Time
	AssignedReviewerIDs []string
}

func (prDB *PullRequestDB) ToDomain() *PullRequest {
	var pullRequest PullRequest
	pullRequest.ID = prDB.ID
	pullRequest.Name = prDB.Name
	pullRequest.AuthorID = prDB.AuthorID
	pullRequest.Status = prDB.Status
	pullRequest.CreatedAt = prDB.CreatedAt
	pullRequest.AssignedReviewerIDs = prDB.AssignedReviewerIDs
	if prDB.MergedAt.Valid {
		pullRequest.MergedAt = &prDB.MergedAt.Time
	}
	return &pullRequest
}
