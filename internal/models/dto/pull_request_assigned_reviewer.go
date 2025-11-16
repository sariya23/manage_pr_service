package dto

type PullRequestAssignedReviewer struct {
	UserID             string
	CountPullRequestID int
}

func ToMapPullRequestAssignedReviewer(prs []PullRequestAssignedReviewer) map[string]int {
	m := make(map[string]int, len(prs))

	for _, pr := range prs {
		m[pr.UserID] = pr.CountPullRequestID
	}
	return m
}
