//go:build integrations

package teams

import (
	"context"
	"net/http"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/sariya23/manage_pr_service/tests/checkers"
	httpclient "github.com/sariya23/manage_pr_service/tests/clients/http"
	"github.com/sariya23/manage_pr_service/tests/factory"
	"github.com/sariya23/manage_pr_service/tests/helpers/random"
	"github.com/sariya23/manage_pr_service/tests/models"
	"github.com/stretchr/testify/require"
)

// TestTeamGet тест на ручку /team/get
// Возвращаются участники команды
func TestTeamGet(t *testing.T) {
	ctx := context.Background()
	httpClient := httpclient.NewHTTPClient()
	members := []factory.AddTeamRequestMemberDTO{}
	for range random.RandInt(1, 3) {
		isActive := true
		members = append(members, factory.RandomInitAddTeamRequestMemberDT("", "", &isActive))
	}
	requestCreateTeam := factory.RandomInitAddTeamRequest("", members)
	responseCreateTeam := httpClient.TeamsAdd(requestCreateTeam)
	require.Equal(t, http.StatusOK, responseCreateTeam.StatusCode)

	response := httpClient.TeamGet(requestCreateTeam.TeamName)
	responseDTO := factory.GetTeamRFromHTTPResponseOK(response)
	teamMembersDB := dbT.GetTeamMembersByTeamName(ctx, requestCreateTeam.TeamName)
	usersDB := dbT.GetUsersFromDB(ctx, models.TeamMemberUserIDs(teamMembersDB))
	checkers.CheckGetTeamResponse(t, responseDTO, teamMembersDB, usersDB)
}

// TestTeamGet_NonexistentTeam тест на ручку /team/get
// Команда не найдена
func TestTeamGet_NonexistentTeam(t *testing.T) {
	httpClient := httpclient.NewHTTPClient()
	response := httpClient.TeamGet(gofakeit.LetterN(10))
	require.Equal(t, http.StatusNotFound, response.StatusCode)
}
