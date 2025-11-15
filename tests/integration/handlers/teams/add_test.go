package teams

import (
	"context"
	"net/http"
	"testing"

	checkers_team "github.com/sariya23/manage_pr_service/tests/checkers/team"
	httpcleint "github.com/sariya23/manage_pr_service/tests/clients/http"
	"github.com/sariya23/manage_pr_service/tests/factory"
	factory_teams "github.com/sariya23/manage_pr_service/tests/factory/teams"
	"github.com/sariya23/manage_pr_service/tests/helpers/random"
	"github.com/stretchr/testify/require"
)

func TestAddTeam_NewTeamNewUsers(t *testing.T) {
	ctx := context.Background()
	httpClient := httpcleint.NewHTTPClient()
	dbT.SetUp(ctx, t, tables...)
	defer dbT.TearDown(t)
	members := []factory_teams.AddTeamRequestMemberDTO{}
	for range random.RandInt(1, 3) {
		isActive := true
		members = append(members, factory_teams.RandomInitAddTeamRequestMemberDT("", "", &isActive))
	}
	request := factory_teams.RandomInitAddTeamRequest("", members)

	response := httpClient.TeamsAdd(request)
	require.Equal(t, http.StatusOK, response.StatusCode)
	responseDTO := factory_teams.FromHTTPResponse(response)
	teamMembersDB := dbT.GetTeamMembersByTeamName(ctx, request.TeamName)
	usersDB := dbT.GetUsersFromDB(ctx, factory.TeamMemberUserIDs(teamMembersDB))
	checkers_team.CheckAddTeamResponse(t, responseDTO, teamMembersDB, usersDB)
}

func TestAddTeam_AddUsersToTeam(t *testing.T) {
	ctx := context.Background()
	httpClient := httpcleint.NewHTTPClient()
	dbT.SetUp(ctx, t, tables...)
	defer dbT.TearDown(t)
	// Предварительно создаем команду с юзерами
	nUsers := random.RandInt(1, 3)
	membersInit := make([]factory_teams.AddTeamRequestMemberDTO, 0, nUsers)
	for range nUsers {
		isActive := true
		membersInit = append(membersInit, factory_teams.RandomInitAddTeamRequestMemberDT("", "", &isActive))
	}
	requestCreateTeam := factory_teams.RandomInitAddTeamRequest("", membersInit)
	responseCreateTeam := httpClient.TeamsAdd(requestCreateTeam)
	require.Equal(t, http.StatusOK, responseCreateTeam.StatusCode)

	nUsers = random.RandInt(1, 3)
	newMembers := make([]factory_teams.AddTeamRequestMemberDTO, 0, nUsers)
	for range nUsers {
		isActive := true
		membersInit = append(membersInit, factory_teams.RandomInitAddTeamRequestMemberDT("", "", &isActive))
	}

	requestAddTeam := factory_teams.RandomInitAddTeamRequest("", newMembers)
	responseAddTeam := httpClient.TeamsAdd(requestAddTeam)
	require.Equal(t, http.StatusOK, responseAddTeam.StatusCode)
	responseDTO := factory_teams.FromHTTPResponse(responseAddTeam)
	teamMembersDB := dbT.GetTeamMembersByTeamName(ctx, requestAddTeam.TeamName)
	usersDB := dbT.GetUsersFromDB(ctx, factory.TeamMemberUserIDs(teamMembersDB))
	checkers_team.CheckAddTeamResponse(t, responseDTO, teamMembersDB, usersDB)
}
