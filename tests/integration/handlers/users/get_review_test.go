//go:build integrations

package users

import (
	"context"
	"net/http"
	"sort"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	httpcleint "github.com/sariya23/manage_pr_service/tests/clients/http"
	"github.com/sariya23/manage_pr_service/tests/factory"
	factory_pull_request "github.com/sariya23/manage_pr_service/tests/factory/pull_request"
	factory_teams "github.com/sariya23/manage_pr_service/tests/factory/teams"
	users_factory "github.com/sariya23/manage_pr_service/tests/factory/users"
	"github.com/sariya23/manage_pr_service/tests/helpers"
	"github.com/sariya23/manage_pr_service/tests/helpers/random"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestUsersGetReview тест на ручку /api/users/getReview
// Успешное получение ревью
func TestUsersGetReview(t *testing.T) {
	ctx := context.Background()
	dbT.SetUp(ctx, t, tables...)
	httpClient := httpcleint.NewHTTPClient()
	nUsers := 2
	members := make([]factory_teams.AddTeamRequestMemberDTO, 0, nUsers)
	for range nUsers {
		isActive := true
		members = append(members, factory_teams.RandomInitAddTeamRequestMemberDT("", "", &isActive))
	}
	requestCreateTeam := factory_teams.RandomInitAddTeamRequest("", members)
	responseCreateTeam := httpClient.TeamsAdd(requestCreateTeam)
	require.Equal(t, http.StatusOK, responseCreateTeam.StatusCode)
	teamDB := dbT.GetTeamMembersByTeamName(ctx, requestCreateTeam.TeamName)
	teamMemberIDs := factory.TeamMemberUserIDs(teamDB)
	authorID := random.Choice(teamMemberIDs)
	memberID := helpers.Filter(teamMemberIDs, authorID)[0]
	requestCreatePR1 := factory_pull_request.PullRequestCreateRequest{}
	requestCreatePR1.RadnomInit("", "", authorID)
	response := httpClient.PullRequestCreate(requestCreatePR1)
	require.Equal(t, http.StatusOK, response.StatusCode)
	requestCreatePR2 := factory_pull_request.PullRequestCreateRequest{}
	requestCreatePR2.RadnomInit("", "", authorID)
	response = httpClient.PullRequestCreate(requestCreatePR2)
	require.Equal(t, http.StatusOK, response.StatusCode)

	responseGet := httpClient.UsersGetReview(memberID)
	require.Equal(t, http.StatusOK, responseGet.StatusCode)
	responseDTO := users_factory.GetReviewFromHTTPResponseOK(responseGet)

	userPRsDB := dbT.GetReviewerPullRequests(ctx, memberID)
	assert.Equal(t, len(userPRsDB), len(responseDTO.PullRequests))
	assert.Equal(t, memberID, responseDTO.UserID)

	sort.Slice(userPRsDB, func(i, j int) bool {
		return userPRsDB[i].ID < userPRsDB[j].ID
	})

	sort.Slice(responseDTO.PullRequests, func(i, j int) bool {
		return responseDTO.PullRequests[i].PullRequestID < responseDTO.PullRequests[j].PullRequestID
	})

	for i := 0; i < len(userPRsDB); i++ {
		assert.Equal(t, userPRsDB[i].ID, responseDTO.PullRequests[i].PullRequestID)
		assert.Equal(t, userPRsDB[i].Name, responseDTO.PullRequests[i].PullRequestName)
		assert.Equal(t, userPRsDB[i].AuthorID, responseDTO.PullRequests[i].AuthorID)
		assert.Equal(t, userPRsDB[i].Status, responseDTO.PullRequests[i].Status)
	}
}

// TestUsersGetReview тест на ручку /api/users/getReview
// Несуществующий юзер
func TestUsersGetReview_NonexistentUser(t *testing.T) {
	httpClient := httpcleint.NewHTTPClient()
	responseGet := httpClient.UsersGetReview(gofakeit.LetterN(8))
	require.Equal(t, http.StatusNotFound, responseGet.StatusCode)
}
