package domain

import "time"

type PullRequest struct {
	ID                  string
	Name                string
	AuthorID            string
	Status              string
	MergedAt            *time.Time
	CreatedAt           time.Time
	AssignedReviewerIDs []string
}
