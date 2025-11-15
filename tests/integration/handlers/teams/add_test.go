package teams

import (
	"context"
	"net/http"
	"testing"

	httpcleint "github.com/sariya23/manage_pr_service/tests/clients/http"
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

	reponseDTO := factory_teams.FromHTTPResponse(response)
	expectedDB := dbT.GetTeamMembersByTeamName(ctx, request.TeamName)

	assert.Equal(t, expectedDB[0].TeamName, reponseDTO.Team.TeamName)
}
