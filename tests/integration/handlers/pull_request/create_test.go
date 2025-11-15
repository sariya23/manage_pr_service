//go:build integrations

package pull_request

import (
	"context"
	"net/http"
	"testing"

	checkers_pull_request "github.com/sariya23/manage_pr_service/tests/checkers/pull_request"
	httpcleint "github.com/sariya23/manage_pr_service/tests/clients/http"
	factory_pull_request "github.com/sariya23/manage_pr_service/tests/factory/pull_request"
	factory_teams "github.com/sariya23/manage_pr_service/tests/factory/teams"
	"github.com/sariya23/manage_pr_service/tests/helpers/random"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

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
