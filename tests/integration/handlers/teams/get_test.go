//go:build integrations

package teams

import (
	"context"
	"net/http"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	checkers_team "github.com/sariya23/manage_pr_service/tests/checkers/team"
	httpcleint "github.com/sariya23/manage_pr_service/tests/clients/http"
	"github.com/sariya23/manage_pr_service/tests/factory"
	factory_teams "github.com/sariya23/manage_pr_service/tests/factory/teams"
	"github.com/sariya23/manage_pr_service/tests/helpers/random"
	"github.com/stretchr/testify/require"
)

// TestTeamGet тест на ручку /api/team/get
// Возвращаются участники команды
func TestTeamGet(t *testing.T) {
	ctx := context.Background()
	httpClient := httpcleint.NewHTTPClient()
	members := []factory_teams.AddTeamRequestMemberDTO{}
	for range random.RandInt(1, 3) {
		isActive := true
		members = append(members, factory_teams.RandomInitAddTeamRequestMemberDT("", "", &isActive))
	}
	requestCreateTeam := factory_teams.RandomInitAddTeamRequest("", members)
	responseCreateTeam := httpClient.TeamsAdd(requestCreateTeam)
	require.Equal(t, http.StatusOK, responseCreateTeam.StatusCode)

	response := httpClient.TeamGet(requestCreateTeam.TeamName)
	responseDTO := factory_teams.GetTeamRFromHTTPResponseOK(response)
	teamMembersDB := dbT.GetTeamMembersByTeamName(ctx, requestCreateTeam.TeamName)
	usersDB := dbT.GetUsersFromDB(ctx, factory.TeamMemberUserIDs(teamMembersDB))
	checkers_team.CheckGetTeamResponse(t, responseDTO, teamMembersDB, usersDB)
}

// TestTeamGet_NonexistentTeam тест на ручку /api/team/get
func TestTeamGet_NonexistentTeam(t *testing.T) {
	httpClient := httpcleint.NewHTTPClient()
	response := httpClient.TeamGet(gofakeit.LetterN(10))
	require.Equal(t, http.StatusNotFound, response.StatusCode)
}
