//go:build integrations

package pull_request

import (
	"context"
	"net/http"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/sariya23/manage_pr_service/tests/checkers"
	httpclient "github.com/sariya23/manage_pr_service/tests/clients/http"
	"github.com/sariya23/manage_pr_service/tests/factory"
	"github.com/sariya23/manage_pr_service/tests/helpers/random"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestPullRequestCreate тест ручки /pullRequest/create
// Успешное создание PullRequest
func TestPullRequestCreate(t *testing.T) {
	ctx := context.Background()
	httpClient := httpclient.NewHTTPClient()
	nUsers := random.RandInt(1, 3)
	members := make([]factory.AddTeamRequestMemberDTO, 0, nUsers)
	for range nUsers {
		isActive := true
		members = append(members, factory.RandomInitAddTeamRequestMemberDT("", "", &isActive))
	}
	requestCreate := factory.RandomInitAddTeamRequest("", members)
	responseCreate := httpClient.TeamsAdd(requestCreate)
	require.Equal(t, http.StatusOK, responseCreate.StatusCode)

	request := factory.PullRequestCreateRequest{}
	request.RadnomInit("", "", random.Choice(members).UserID)

	response := httpClient.PullRequestCreate(request)
	require.Equal(t, http.StatusOK, response.StatusCode)
	responseDTO := factory.PullRequestCreateFromHTTPResponseOK(response)
	pullRequestDB := dbT.GetPullRequest(ctx, request.PullRequestID)

	assert.Equal(t, request.PullRequestID, pullRequestDB.ID)
	assert.Equal(t, request.PullRequestName, pullRequestDB.Name)
	assert.Equal(t, request.AuthorID, pullRequestDB.AuthorID)

	checkers.CheckPullRequestCreateResponse(t, responseDTO, *pullRequestDB)
}

// TestPullRequestCreate_AuthorNotFound тест ручки /pullRequest/create
// При попытке создать PR несуществующим пользователем, вернется ошибка
func TestPullRequestCreate_AuthorNotFound(t *testing.T) {
	httpClient := httpclient.NewHTTPClient()
	request := factory.PullRequestCreateRequest{}
	request.RadnomInit("", "", "")
	response := httpClient.PullRequestCreate(request)
	require.Equal(t, http.StatusNotFound, response.StatusCode)
}

// TestPullRequestCreate_AlreadyExists тест ручки /pullRequest/create
// При попытке создать PR с уже существующим айдишником, вернется ошибка
func TestPullRequestCreate_AlreadyExists(t *testing.T) {
	httpClient := httpclient.NewHTTPClient()
	nUsers := random.RandInt(1, 3)
	members := make([]factory.AddTeamRequestMemberDTO, 0, nUsers)
	for range nUsers {
		isActive := true
		members = append(members, factory.RandomInitAddTeamRequestMemberDT("", "", &isActive))
	}
	requestCreate := factory.RandomInitAddTeamRequest("", members)
	responseCreate := httpClient.TeamsAdd(requestCreate)
	require.Equal(t, http.StatusOK, responseCreate.StatusCode)

	requestCreatePR := factory.PullRequestCreateRequest{}
	requestCreatePR.RadnomInit("", "", random.Choice(members).UserID)
	responseCreatePR := httpClient.PullRequestCreate(requestCreatePR)
	require.Equal(t, http.StatusOK, responseCreatePR.StatusCode)

	responseAlreadyExists := httpClient.PullRequestCreate(requestCreatePR)
	require.Equal(t, http.StatusConflict, responseAlreadyExists.StatusCode)
}

// TestPullRequestCreate_ValidationError тест ручки /pullRequest/create
// Ошибки валидации
func TestPullRequestCreate_ValidationError(t *testing.T) {
	cases := []struct {
		name    string
		request factory.PullRequestCreateRequest
	}{
		{
			name: "empty pull request id",
			request: factory.PullRequestCreateRequest{
				PullRequestName: gofakeit.LetterN(32),
				AuthorID:        gofakeit.LetterN(32)},
		},
		{
			name: "empty author id",
			request: factory.PullRequestCreateRequest{
				PullRequestName: gofakeit.LetterN(32),
				PullRequestID:   gofakeit.LetterN(32)},
		},
		{
			name: "empty pull request name",
			request: factory.PullRequestCreateRequest{
				AuthorID:      gofakeit.LetterN(32),
				PullRequestID: gofakeit.LetterN(32)},
		},
	}
	httpClient := httpclient.NewHTTPClient()
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			response := httpClient.PullRequestCreate(c.request)
			require.Equal(t, http.StatusBadRequest, response.StatusCode)
		})
	}
}
