//go:build integrations

package pull_request

import (
	"context"
	"net/http"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/sariya23/manage_pr_service/tests/checkers"
	httpcleint "github.com/sariya23/manage_pr_service/tests/clients/http"
	"github.com/sariya23/manage_pr_service/tests/factory"
	"github.com/sariya23/manage_pr_service/tests/helpers/random"
	"github.com/stretchr/testify/require"
)

// TestPullRequestMerge тест на ручку /api/pullRequest/merge
// УСпешный merge
func TestPullRequestMerge(t *testing.T) {
	ctx := context.Background()
	dbT.SetUp(ctx, t, tables...)
	httpClient := httpcleint.NewHTTPClient()
	nUsers := random.RandInt(1, 3)
	members := make([]factory.AddTeamRequestMemberDTO, 0, nUsers)
	for range nUsers {
		isActive := true
		members = append(members, factory.RandomInitAddTeamRequestMemberDT("", "", &isActive))
	}
	requestCreateTeam := factory.RandomInitAddTeamRequest("", members)
	responseCreateTeam := httpClient.TeamsAdd(requestCreateTeam)
	require.Equal(t, http.StatusOK, responseCreateTeam.StatusCode)
	requestCreatePR := factory.PullRequestCreateRequest{}
	requestCreatePR.RadnomInit("", "", random.Choice(members).UserID)
	responseCreatePR := httpClient.PullRequestCreate(requestCreatePR)
	require.Equal(t, http.StatusOK, responseCreatePR.StatusCode)

	requestMerge := factory.PullRequestMergeRequest{PullRequestID: requestCreatePR.PullRequestID}
	responseMerge := httpClient.PullRequestMerge(requestMerge)
	require.Equal(t, http.StatusOK, responseMerge.StatusCode)
	responseDTO := factory.PullRequestMergeFromHTTPResponseOK(responseMerge)

	pullRequestDB := dbT.GetPullRequest(ctx, requestCreatePR.PullRequestID)
	checkers.CheckPullRequestMergeResponse(t, responseDTO, *pullRequestDB)
}

// TestPullRequestMerge тест на ручку /api/pullRequest/merge
// Ошибка при попытке вмержить несуществующий PR
func TestPullRequestMerge_NotFound(t *testing.T) {
	httpClient := httpcleint.NewHTTPClient()
	requestMerge := factory.PullRequestMergeRequest{PullRequestID: gofakeit.LetterN(8)}
	responseMerge := httpClient.PullRequestMerge(requestMerge)
	require.Equal(t, http.StatusNotFound, responseMerge.StatusCode)
}

// TestPullRequestMerge_EmptyPullRequestID тест на ручку /api/pullRequest/merge
// Ошибка при передаче пустого айди
func TestPullRequestMerge_EmptyPullRequestID(t *testing.T) {
	httpClient := httpcleint.NewHTTPClient()
	requestMerge := factory.PullRequestMergeRequest{}
	responseMerge := httpClient.PullRequestMerge(requestMerge)
	require.Equal(t, http.StatusBadRequest, responseMerge.StatusCode)
}
