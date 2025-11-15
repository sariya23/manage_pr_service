//go:build integrations

package checkers_team

import (
	"testing"

	"github.com/sariya23/manage_pr_service/tests/factory"
	factory_teams "github.com/sariya23/manage_pr_service/tests/factory/teams"
	"github.com/stretchr/testify/assert"
)

func CheckAddTeamResponse(t *testing.T, responseDTO factory_teams.AddTeamResponse, teamMembersDB []factory.TeamMember,
	usersDB []factory.User) {
	factory.TeamMemberSortByUserID(teamMembersDB)
	factory.UserSortByUserID(usersDB)
	factory_teams.AddTeamResponseMemberSortByUserID(responseDTO.Team.Members)
	assert.Equal(t, len(teamMembersDB), len(responseDTO.Team.Members))
	for i := 0; i < len(teamMembersDB); i++ {
		assert.Equal(t, teamMembersDB[i].TeamName, responseDTO.Team.TeamName)
		assert.Equal(t, usersDB[i].UserID, responseDTO.Team.Members[i].UserID)
		assert.Equal(t, usersDB[i].Username, responseDTO.Team.Members[i].Username)
		assert.Equal(t, usersDB[i].IsActive, responseDTO.Team.Members[i].IsActive)
	}
}
