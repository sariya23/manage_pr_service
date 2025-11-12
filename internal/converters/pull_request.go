package converters

import (
	"strconv"

	api "github.com/sariya23/manage_pr_service/internal/generated"
	"github.com/sariya23/manage_pr_service/internal/models"
)

func DomainPullRequestToGetReviewResponse(domainPR models.PullRequest) api.PullRequestShort {
	var res api.PullRequestShort
	res.PullRequestId = strconv.Itoa(int(domainPR.PullRequestID))
	res.AuthorId = strconv.Itoa(int(domainPR.AuthorID))
	res.PullRequestName = domainPR.Name
	res.Status = api.PullRequestShortStatus(domainPR.Status)
	return res
}

func MultiDomainPullRequestToGetReviewResponse(domainPRs []models.PullRequest) []api.PullRequestShort {
	res := make([]api.PullRequestShort, 0, len(domainPRs))
	for _, pr := range domainPRs {
		res = append(res, DomainPullRequestToGetReviewResponse(pr))
	}
	return res
}
