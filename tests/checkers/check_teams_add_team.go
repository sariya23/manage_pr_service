//go:build integrations

package checkers

import (
	"testing"

	"github.com/sariya23/manage_pr_service/tests/factory"
	"github.com/sariya23/manage_pr_service/tests/models"
	"github.com/stretchr/testify/assert"
)

func CheckAddTeamResponse(t *testing.T, responseDTO factory.AddTeamResponse, teamMembersDB []models.TeamMember,
	usersDB []models.User) {
	models.TeamMemberSortByUserID(teamMembersDB)
	models.UserSortByUserID(usersDB)
	factory.AddTeamResponseMemberSortByUserID(responseDTO.Team.Members)
	assert.Equal(t, len(teamMembersDB), len(responseDTO.Team.Members))
	for i := 0; i < len(teamMembersDB); i++ {
		assert.Equal(t, teamMembersDB[i].TeamName, responseDTO.Team.TeamName)
		assert.Equal(t, usersDB[i].UserID, responseDTO.Team.Members[i].UserID)
		assert.Equal(t, usersDB[i].Username, responseDTO.Team.Members[i].Username)
		assert.Equal(t, usersDB[i].IsActive, responseDTO.Team.Members[i].IsActive)
	}
}
