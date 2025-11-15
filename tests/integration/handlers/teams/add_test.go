package teams

import (
	"context"
	"net/http"
	"testing"

	httpcleint "github.com/sariya23/manage_pr_service/tests/clients/http"
	"github.com/sariya23/manage_pr_service/tests/factory"
	factory_teams "github.com/sariya23/manage_pr_service/tests/factory/teams"
	"github.com/sariya23/manage_pr_service/tests/helpers/random"
	"github.com/stretchr/testify/assert"
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
	factory.TeamMemberSortByUserID(teamMembersDB)
	usersDB := dbT.GetUsersFromDB(ctx, factory.TeamMemberUserIDs(teamMembersDB))
	factory.UserSortByUserID(usersDB)
	assert.Equal(t, len(teamMembersDB), len(responseDTO.Team.Members))
	for i := 0; i < len(teamMembersDB); i++ {
		assert.Equal(t, teamMembersDB[i].TeamName, responseDTO.Team.TeamName)

		assert.Equal(t, usersDB[i].UserID, responseDTO.Team.Members[i].UserID)
		assert.Equal(t, usersDB[i].Username, responseDTO.Team.Members[i].Username)
		assert.Equal(t, usersDB[i].IsActive, responseDTO.Team.Members[i].IsActive)
	}
}
