package converters

import (
	api "github.com/sariya23/manage_pr_service/internal/generated"
	"github.com/sariya23/manage_pr_service/internal/models/domain"
)

func DomainPullRequestToGetReviewResponse(domainPR domain.PullRequest) api.PullRequestShort {
	var res api.PullRequestShort
	res.PullRequestId = domainPR.ID
	res.AuthorId = domainPR.AuthorID
	res.PullRequestName = domainPR.Name
	res.Status = api.PullRequestShortStatus(domainPR.Status)
	return res
}

func MultiDomainPullRequestToGetReviewResponse(domainPRs []domain.PullRequest) []api.PullRequestShort {
	res := make([]api.PullRequestShort, 0, len(domainPRs))
	for _, pr := range domainPRs {
		res = append(res, DomainPullRequestToGetReviewResponse(pr))
	}
	return res
}

func DomainPullRequestToCreatePullRequestResponse(domainPR domain.PullRequest) api.PullRequest {
	var pr api.PullRequest
	pr.PullRequestId = domainPR.ID
	pr.AuthorId = domainPR.AuthorID
	pr.PullRequestName = domainPR.Name
	pr.Status = api.PullRequestStatus(domainPR.Status)
	pr.MergedAt = &domainPR.MergedAt
	pr.CreatedAt = &domainPR.CreatedAt
	return pr
}
