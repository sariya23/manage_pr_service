//go:build integrations

package pull_request

import (
	"context"
	"net/http"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	checkers_pull_request "github.com/sariya23/manage_pr_service/tests/checkers/pull_request"
	httpcleint "github.com/sariya23/manage_pr_service/tests/clients/http"
	factory_pull_request "github.com/sariya23/manage_pr_service/tests/factory/pull_request"
	factory_teams "github.com/sariya23/manage_pr_service/tests/factory/teams"
	"github.com/sariya23/manage_pr_service/tests/helpers/random"
	"github.com/stretchr/testify/require"
)

// TestPullRequestMerge тест на ручку /api/pullRequest/merge
// УСпешный merge
func TestPullRequestMerge(t *testing.T) {
	ctx := context.Background()
	httpClient := httpcleint.NewHTTPClient()
	nUsers := random.RandInt(1, 3)
	members := make([]factory_teams.AddTeamRequestMemberDTO, 0, nUsers)
	for range nUsers {
		isActive := true
		members = append(members, factory_teams.RandomInitAddTeamRequestMemberDT("", "", &isActive))
	}
	requestCreateTeam := factory_teams.RandomInitAddTeamRequest("", members)
	responseCreateTeam := httpClient.TeamsAdd(requestCreateTeam)
	require.Equal(t, http.StatusOK, responseCreateTeam.StatusCode)
	requestCreatePR := factory_pull_request.PullRequestCreateRequest{}
	requestCreatePR.RadnomInit("", "", random.Choice(members).UserID)
	responseCreatePR := httpClient.PullRequestCreate(requestCreatePR)
	require.Equal(t, http.StatusOK, responseCreatePR.StatusCode)

	requestMerge := factory_pull_request.PullRequestMergeRequest{PullRequestID: requestCreatePR.PullRequestID}
	responseMerge := httpClient.PullRequestMerge(requestMerge)
	require.Equal(t, http.StatusOK, responseMerge.StatusCode)
	responseDTO := factory_pull_request.PullRequestMergeFromHTTPResponseOK(responseMerge)

	pullRequestDB := dbT.GetPullRequest(ctx, requestCreatePR.PullRequestID)
	checkers_pull_request.CheckPullRequestMergeResponse(t, responseDTO, *pullRequestDB)
}

// TestPullRequestMerge тест на ручку /api/pullRequest/merge
// Ошибка при попытке вмержить несуществующий PR
func TestPullRequestMerge_NotFound(t *testing.T) {
	httpClient := httpcleint.NewHTTPClient()
	requestMerge := factory_pull_request.PullRequestMergeRequest{PullRequestID: gofakeit.LetterN(8)}
	responseMerge := httpClient.PullRequestMerge(requestMerge)
	require.Equal(t, http.StatusNotFound, responseMerge.StatusCode)
}

// TestPullRequestMerge_EmptyPullRequestID тест на ручку /api/pullRequest/merge
// Ошибка при передаче пустого айди
func TestPullRequestMerge_EmptyPullRequestID(t *testing.T) {
	httpClient := httpcleint.NewHTTPClient()
	requestMerge := factory_pull_request.PullRequestMergeRequest{}
	responseMerge := httpClient.PullRequestMerge(requestMerge)
	require.Equal(t, http.StatusBadRequest, responseMerge.StatusCode)
}
