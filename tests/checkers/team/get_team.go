//go:build integrations

package checkers_team

import (
	"testing"

	"github.com/sariya23/manage_pr_service/tests/factory"
	factory_teams "github.com/sariya23/manage_pr_service/tests/factory/teams"
	"github.com/stretchr/testify/assert"
)

func CheckGetTeamResponse(t *testing.T, responseDTO factory_teams.GetTeamResponse, teamMembersDB []factory.TeamMember,
	usersDB []factory.User) {
	factory.TeamMemberSortByUserID(teamMembersDB)
	factory.UserSortByUserID(usersDB)
	factory_teams.GetTeamResponseMemberSortByUserID(responseDTO.Members)
	assert.Equal(t, len(teamMembersDB), len(responseDTO.Members))
	for i := 0; i < len(teamMembersDB); i++ {
		assert.Equal(t, teamMembersDB[i].TeamName, responseDTO.TeamName)
		assert.Equal(t, usersDB[i].UserID, responseDTO.Members[i].UserID)
		assert.Equal(t, usersDB[i].Username, responseDTO.Members[i].Username)
		assert.Equal(t, usersDB[i].IsActive, responseDTO.Members[i].IsActive)
	}
}
