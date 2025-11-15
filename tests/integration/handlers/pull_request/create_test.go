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
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestPullRequestCreate тест ручки /api/pullRequest/create
// Успешное создание PullRequest
func TestPullRequestCreate(t *testing.T) {
	ctx := context.Background()
	httpClient := httpcleint.NewHTTPClient()
	nUsers := random.RandInt(1, 3)
	members := make([]factory_teams.AddTeamRequestMemberDTO, 0, nUsers)
	for range nUsers {
		isActive := true
		members = append(members, factory_teams.RandomInitAddTeamRequestMemberDT("", "", &isActive))
	}
	requestCreate := factory_teams.RandomInitAddTeamRequest("", members)
	responseCreate := httpClient.TeamsAdd(requestCreate)
	require.Equal(t, http.StatusOK, responseCreate.StatusCode)

	request := factory_pull_request.PullRequestCreateRequest{}
	request.RadnomInit("", "", random.Choice(members).UserID)

	response := httpClient.PullRequestCreate(request)
	require.Equal(t, http.StatusOK, response.StatusCode)
	responseDTO := factory_pull_request.PullRequestCreateFromHTTPResponseOK(response)
	pullRequestDB := dbT.GetPullRequest(ctx, request.PullRequestID)

	assert.Equal(t, request.PullRequestID, pullRequestDB.ID)
	assert.Equal(t, request.PullRequestName, pullRequestDB.Name)
	assert.Equal(t, request.AuthorID, pullRequestDB.AuthorID)

	checkers_pull_request.CheckPullRequestCreateResponse(t, responseDTO, *pullRequestDB)
}

// TestPullRequestCreate_AuthorNotFound тест ручки /api/pullRequest/create
// При попытке создать PR несуществующим пользователем, вернется ошибка
func TestPullRequestCreate_AuthorNotFound(t *testing.T) {
	httpClient := httpcleint.NewHTTPClient()
	request := factory_pull_request.PullRequestCreateRequest{}
	request.RadnomInit("", "", "")
	response := httpClient.PullRequestCreate(request)
	require.Equal(t, http.StatusNotFound, response.StatusCode)
}

// TestPullRequestCreate_AlreadyExists тест ручки /api/pullRequest/create
// При попытке создать PR с уже существующим айдишником, вернется ошибка
func TestPullRequestCreate_AlreadyExists(t *testing.T) {
	httpClient := httpcleint.NewHTTPClient()
	nUsers := random.RandInt(1, 3)
	members := make([]factory_teams.AddTeamRequestMemberDTO, 0, nUsers)
	for range nUsers {
		isActive := true
		members = append(members, factory_teams.RandomInitAddTeamRequestMemberDT("", "", &isActive))
	}
	requestCreate := factory_teams.RandomInitAddTeamRequest("", members)
	responseCreate := httpClient.TeamsAdd(requestCreate)
	require.Equal(t, http.StatusOK, responseCreate.StatusCode)

	requestCreatePR := factory_pull_request.PullRequestCreateRequest{}
	requestCreatePR.RadnomInit("", "", random.Choice(members).UserID)
	responseCreatePR := httpClient.PullRequestCreate(requestCreatePR)
	require.Equal(t, http.StatusOK, responseCreatePR.StatusCode)

	responseAlreadyExists := httpClient.PullRequestCreate(requestCreatePR)
	require.Equal(t, http.StatusConflict, responseAlreadyExists.StatusCode)
}

// TestPullRequestCreate_ValidationError тест ручки /api/pullRequest/create
// Ошибки валидации
func TestPullRequestCreate_ValidationError(t *testing.T) {
	cases := []struct {
		name    string
		request factory_pull_request.PullRequestCreateRequest
	}{
		{
			name: "empty pull request id",
			request: factory_pull_request.PullRequestCreateRequest{
				PullRequestName: gofakeit.LetterN(8),
				AuthorID:        gofakeit.LetterN(8)},
		},
		{
			name: "empty author id",
			request: factory_pull_request.PullRequestCreateRequest{
				PullRequestName: gofakeit.LetterN(8),
				PullRequestID:   gofakeit.LetterN(8)},
		},
		{
			name: "empty pull request name",
			request: factory_pull_request.PullRequestCreateRequest{
				AuthorID:      gofakeit.LetterN(8),
				PullRequestID: gofakeit.LetterN(8)},
		},
	}
	httpClient := httpcleint.NewHTTPClient()
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			response := httpClient.PullRequestCreate(c.request)
			require.Equal(t, http.StatusBadRequest, response.StatusCode)
		})
	}
}
