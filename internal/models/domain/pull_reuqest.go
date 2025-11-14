package domain

import "time"

type PullRequestStatus string

const (
	PULL_REQUEST_STATUS_OPEN   PullRequestStatus = "OPEN"
	PULL_REQUEST_STATUS_MERGED PullRequestStatus = "MERGED"
)

type PullRequest struct {
	ID        string
	Name      string
	AuthorID  string
	Status    PullRequestStatus
	MergedAt  time.Time
	CreatedAt time.Time
}
