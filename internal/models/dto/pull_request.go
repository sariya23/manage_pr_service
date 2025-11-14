package dto

import api "github.com/sariya23/manage_pr_service/internal/generated"

type CreatePullRequestDTO struct {
	ID       string
	Name     string
	AuthorID string
}

func FromCreatePullRequestHTTP(req api.PostPullRequestCreateJSONRequestBody) CreatePullRequestDTO {
	return CreatePullRequestDTO{
		ID:       req.PullRequestId,
		Name:     req.PullRequestName,
		AuthorID: req.AuthorId,
	}
}
