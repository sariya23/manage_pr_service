//go:build integrations

package pull_request

import (
	"context"
	"net/http"
	"slices"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/sariya23/manage_pr_service/tests/checkers"
	httpcleint "github.com/sariya23/manage_pr_service/tests/clients/http"
	"github.com/sariya23/manage_pr_service/tests/factory"
	"github.com/sariya23/manage_pr_service/tests/helpers"
	"github.com/sariya23/manage_pr_service/tests/helpers/random"
	"github.com/sariya23/manage_pr_service/tests/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestPullRequestReassign тест ручки /api/pullRequest/reassign
// Успешный reassign
func TestPullRequestReassign(t *testing.T) {
	ctx := context.Background()
	dbT.SetUp(ctx, t, tables...)
	httpClient := httpcleint.NewHTTPClient()
	nUsers := random.RandInt(3, 6)
	members := make([]factory.AddTeamRequestMemberDTO, 0, nUsers)
	for range nUsers {
		isActive := true
		members = append(members, factory.RandomInitAddTeamRequestMemberDT("", "", &isActive))
	}
	requestCreateTeam := factory.RandomInitAddTeamRequest("", members)
	responseCreateTeam := httpClient.TeamsAdd(requestCreateTeam)
	require.Equal(t, http.StatusOK, responseCreateTeam.StatusCode)
	requestCreatePR := factory.PullRequestCreateRequest{}
	authorID := random.Choice(members).UserID
	requestCreatePR.RadnomInit("", "", authorID)
	responseCreatePR := httpClient.PullRequestCreate(requestCreatePR)
	require.Equal(t, http.StatusOK, responseCreatePR.StatusCode)

	prReviewers := dbT.GetPullRequest(ctx, requestCreatePR.PullRequestID).AssignedReviewerIDs
	requestReassign := factory.PullRequestReassignRequest{
		PullRequestID: requestCreatePR.PullRequestID,
		OldUserID:     random.Choice(helpers.Filter(prReviewers, authorID)),
	}
	responseReassign := httpClient.PullRequestReassign(requestReassign)
	require.Equal(t, http.StatusOK, responseReassign.StatusCode)
	responseDTO := factory.PullRequestReassignFromHTTPResponseOK(responseReassign)
	prDB := dbT.GetPullRequest(ctx, requestCreatePR.PullRequestID)

	assert.False(t, slices.Contains(prDB.AssignedReviewerIDs, authorID))
	checkers.CheckPullRequestReassignResponse(t, responseDTO, *prDB)
}

// TestPullRequestReassign_NonexistentPullRequest тест ручки /api/pullRequest/reassign
// Несуществующий PR
func TestPullRequestReassign_NonexistentPullRequest(t *testing.T) {
	httpClient := httpcleint.NewHTTPClient()
	requestReassign := factory.PullRequestReassignRequest{
		PullRequestID: gofakeit.LetterN(8),
		OldUserID:     gofakeit.LetterN(8),
	}
	response := httpClient.PullRequestReassign(requestReassign)
	require.Equal(t, http.StatusNotFound, response.StatusCode)
}

// TestPullRequestReassign_PullRequestMerged тест ручки /api/pullRequest/reassign
// Нельзя изменить вмерженный PR
func TestPullRequestReassign_PullRequestMerged(t *testing.T) {
	httpClient := httpcleint.NewHTTPClient()
	ctx := context.Background()
	dbT.SetUp(ctx, t, tables...)
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

	requestReassign := factory.PullRequestReassignRequest{
		PullRequestID: requestCreatePR.PullRequestID,
		OldUserID:     gofakeit.LetterN(8),
	}
	response := httpClient.PullRequestReassign(requestReassign)
	require.Equal(t, http.StatusConflict, response.StatusCode)
}

// TestPullRequestReassign_NonexistentOldUserID тест ручки /api/pullRequest/reassign
// Несуществующий переназначаемый пользователь
func TestPullRequestReassign_NonexistentOldUserID(t *testing.T) {
	httpClient := httpcleint.NewHTTPClient()
	ctx := context.Background()
	dbT.SetUp(ctx, t, tables...)
	nUsers := random.RandInt(4, 6)
	members := make([]factory.AddTeamRequestMemberDTO, 0, nUsers)
	for range nUsers {
		isActive := true
		members = append(members, factory.RandomInitAddTeamRequestMemberDT("", "", &isActive))
	}
	requestCreateTeam := factory.RandomInitAddTeamRequest("", members)
	responseCreateTeam := httpClient.TeamsAdd(requestCreateTeam)
	require.Equal(t, http.StatusOK, responseCreateTeam.StatusCode)
	requestCreatePR := factory.PullRequestCreateRequest{}
	authorID := random.Choice(members).UserID
	requestCreatePR.RadnomInit("", "", authorID)
	responseCreatePR := httpClient.PullRequestCreate(requestCreatePR)
	require.Equal(t, http.StatusOK, responseCreatePR.StatusCode)

	requestReassign := factory.PullRequestReassignRequest{
		PullRequestID: requestCreatePR.PullRequestID,
		OldUserID:     gofakeit.LetterN(8),
	}
	response := httpClient.PullRequestReassign(requestReassign)
	require.Equal(t, http.StatusNotFound, response.StatusCode)
}

// TestPullRequestReassign_OldUserIsAuthor тест ручки /api/pullRequest/reassign
// Переназначаемый юзер - это автор PR
func TestPullRequestReassign_OldUserIsAuthor(t *testing.T) {
	httpClient := httpcleint.NewHTTPClient()
	ctx := context.Background()
	dbT.SetUp(ctx, t, tables...)
	nUsers := random.RandInt(4, 6)
	members := make([]factory.AddTeamRequestMemberDTO, 0, nUsers)
	for range nUsers {
		isActive := true
		members = append(members, factory.RandomInitAddTeamRequestMemberDT("", "", &isActive))
	}
	requestCreateTeam := factory.RandomInitAddTeamRequest("", members)
	responseCreateTeam := httpClient.TeamsAdd(requestCreateTeam)
	require.Equal(t, http.StatusOK, responseCreateTeam.StatusCode)
	requestCreatePR := factory.PullRequestCreateRequest{}
	authorID := random.Choice(members).UserID
	requestCreatePR.RadnomInit("", "", authorID)
	responseCreatePR := httpClient.PullRequestCreate(requestCreatePR)
	require.Equal(t, http.StatusOK, responseCreatePR.StatusCode)

	requestReassign := factory.PullRequestReassignRequest{
		PullRequestID: requestCreatePR.PullRequestID,
		OldUserID:     requestCreatePR.AuthorID,
	}
	response := httpClient.PullRequestReassign(requestReassign)
	require.Equal(t, http.StatusConflict, response.StatusCode)
}

// TestPullRequestReassign_OldUserInAnotherTeam тест ручки /api/pullRequest/reassign
// Переназначаемый юзер не принадлежит команде ПРа
func TestPullRequestReassign_OldUserInAnotherTeam(t *testing.T) {
	ctx := context.Background()
	dbT.SetUp(ctx, t, tables...)
	httpClient := httpcleint.NewHTTPClient()
	nUsers := random.RandInt(4, 6)
	members1 := make([]factory.AddTeamRequestMemberDTO, 0, nUsers)
	for range nUsers {
		isActive := true
		members1 = append(members1, factory.RandomInitAddTeamRequestMemberDT("", "", &isActive))
	}
	requestCreateTeam1 := factory.RandomInitAddTeamRequest("", members1)
	responseCreateTeam1 := httpClient.TeamsAdd(requestCreateTeam1)
	require.Equal(t, http.StatusOK, responseCreateTeam1.StatusCode)
	members2 := make([]factory.AddTeamRequestMemberDTO, 0, nUsers)
	for range nUsers {
		isActive := true
		members2 = append(members2, factory.RandomInitAddTeamRequestMemberDT("", "", &isActive))
	}
	requestCreateTeam2 := factory.RandomInitAddTeamRequest("", members2)
	responseCreateTeam2 := httpClient.TeamsAdd(requestCreateTeam2)
	require.Equal(t, http.StatusOK, responseCreateTeam2.StatusCode)

	requestCreatePR := factory.PullRequestCreateRequest{}
	authorID := random.Choice(members1).UserID
	requestCreatePR.RadnomInit("", "", authorID)
	responseCreatePR := httpClient.PullRequestCreate(requestCreatePR)
	require.Equal(t, http.StatusOK, responseCreatePR.StatusCode)

	requestReassign := factory.PullRequestReassignRequest{
		PullRequestID: requestCreatePR.PullRequestID,
		OldUserID:     random.Choice(members2).UserID,
	}
	response := httpClient.PullRequestReassign(requestReassign)
	require.Equal(t, http.StatusBadRequest, response.StatusCode)
}

// TestPullRequestReassign_OldUserIsNotReviewer тест ручки /api/pullRequest/reassign
// Переназначаемый юзер не в ревьюверах ПРа
func TestPullRequestReassign_OldUserIsNotReviewer(t *testing.T) {
	ctx := context.Background()
	dbT.SetUp(ctx, t, tables...)
	httpClient := httpcleint.NewHTTPClient()
	nUsers := random.RandInt(4, 6)
	members := make([]factory.AddTeamRequestMemberDTO, 0, nUsers)
	for range nUsers {
		isActive := true
		members = append(members, factory.RandomInitAddTeamRequestMemberDT("", "", &isActive))
	}
	requestCreateTeam := factory.RandomInitAddTeamRequest("", members)
	responseCreateTeam := httpClient.TeamsAdd(requestCreateTeam)
	require.Equal(t, http.StatusOK, responseCreateTeam.StatusCode)
	requestCreatePR := factory.PullRequestCreateRequest{}
	authorID := random.Choice(members).UserID
	requestCreatePR.RadnomInit("", "", authorID)
	responseCreatePR := httpClient.PullRequestCreate(requestCreatePR)
	require.Equal(t, http.StatusOK, responseCreatePR.StatusCode)

	teamDB := dbT.GetTeamMembersByTeamName(ctx, requestCreateTeam.TeamName)
	prDB := dbT.GetPullRequest(ctx, requestCreatePR.PullRequestID)
	membersWithoutAuthor := helpers.Filter(models.TeamMemberUserIDs(teamDB), requestCreatePR.AuthorID)
	requestReassign := factory.PullRequestReassignRequest{
		PullRequestID: requestCreatePR.PullRequestID,
		OldUserID:     random.Choice(helpers.Filters(membersWithoutAuthor, prDB.AssignedReviewerIDs)),
	}
	response := httpClient.PullRequestReassign(requestReassign)
	require.Equal(t, http.StatusConflict, response.StatusCode)
}

// TestPullRequestReassign_NoReviewerCandidates тест ручки /api/pullRequest/reassign
// Нет кандидатов на переназначение
func TestPullRequestReassign_NoReviewerCandidates(t *testing.T) {
	ctx := context.Background()
	dbT.SetUp(ctx, t, tables...)
	httpClient := httpcleint.NewHTTPClient()
	nUsers := 2
	members := make([]factory.AddTeamRequestMemberDTO, 0, nUsers)
	for range nUsers {
		isActive := true
		members = append(members, factory.RandomInitAddTeamRequestMemberDT("", "", &isActive))
	}
	requestCreateTeam := factory.RandomInitAddTeamRequest("", members)
	responseCreateTeam := httpClient.TeamsAdd(requestCreateTeam)
	require.Equal(t, http.StatusOK, responseCreateTeam.StatusCode)
	requestCreatePR := factory.PullRequestCreateRequest{}
	authorID := random.Choice(members).UserID
	requestCreatePR.RadnomInit("", "", authorID)
	responseCreatePR := httpClient.PullRequestCreate(requestCreatePR)
	require.Equal(t, http.StatusOK, responseCreatePR.StatusCode)

	teamDB := dbT.GetTeamMembersByTeamName(ctx, requestCreateTeam.TeamName)
	inActiveUser := random.Choice(helpers.Filter(models.TeamMemberUserIDs(teamDB), authorID))
	requestSetIsActive := factory.SetIsActiveRequest{
		UserID:   inActiveUser,
		IsActive: false,
	}
	responseSetIsActive := httpClient.UsersSetIsActive(requestSetIsActive)
	require.Equal(t, http.StatusOK, responseSetIsActive.StatusCode)

	requestReassign := factory.PullRequestReassignRequest{
		PullRequestID: requestCreatePR.PullRequestID,
		OldUserID:     inActiveUser,
	}
	response := httpClient.PullRequestReassign(requestReassign)
	require.Equal(t, http.StatusConflict, response.StatusCode)
}

// TestPullRequestReassign_ValidationError  тест ручки /api/pullRequest/reassign
// Ошибки валидации
func TestPullRequestReassign_ValidationError(t *testing.T) {
	cases := []struct {
		name    string
		request factory.PullRequestReassignRequest
	}{
		{
			name:    "no pr id",
			request: factory.PullRequestReassignRequest{OldUserID: "zxc"},
		},
		{
			name:    "no user id",
			request: factory.PullRequestReassignRequest{PullRequestID: "zxc"},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			httpClient := httpcleint.NewHTTPClient()
			response := httpClient.PullRequestReassign(c.request)
			require.Equal(t, http.StatusBadRequest, response.StatusCode)
		})
	}
}
