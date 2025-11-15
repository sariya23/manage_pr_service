package checkers

import (
	"slices"
	"sort"
	"testing"

	"github.com/sariya23/manage_pr_service/tests/factory"
	"github.com/sariya23/manage_pr_service/tests/models"
	"github.com/stretchr/testify/assert"
)

func CheckPullRequestCreateResponse(t *testing.T, responseDTO factory.PullRequestCreateResponse,
	pullRequestDB models.PullRequest) {
	assert.Equal(t, pullRequestDB.ID, responseDTO.PR.PullRequestID)
	assert.Equal(t, pullRequestDB.Name, responseDTO.PR.PullRequestName)
	assert.Equal(t, pullRequestDB.AuthorID, responseDTO.PR.AuthorID)
	assert.Equal(t, pullRequestDB.Status, responseDTO.PR.Status)
	sort.Slice(pullRequestDB.AssignedReviewerIDs, func(i, j int) bool {
		return pullRequestDB.AssignedReviewerIDs[i] < pullRequestDB.AssignedReviewerIDs[j]
	})
	sort.Slice(responseDTO.PR.AssignedReviewers, func(i, j int) bool {
		return responseDTO.PR.AssignedReviewers[i] < responseDTO.PR.AssignedReviewers[j]
	})
	assert.Equal(t, pullRequestDB.AssignedReviewerIDs, responseDTO.PR.AssignedReviewers)
	assert.False(t, slices.Contains(pullRequestDB.AssignedReviewerIDs, pullRequestDB.AuthorID))
}
