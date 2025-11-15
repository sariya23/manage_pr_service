package checkers

import (
	"testing"

	"github.com/sariya23/manage_pr_service/tests/factory"
	"github.com/sariya23/manage_pr_service/tests/models"
	"github.com/stretchr/testify/assert"
)

func CheckGetTeamResponse(t *testing.T, responseDTO factory.GetTeamResponse, teamMembersDB []models.TeamMember,
	usersDB []models.User) {
	models.TeamMemberSortByUserID(teamMembersDB)
	models.UserSortByUserID(usersDB)
	factory.GetTeamResponseMemberSortByUserID(responseDTO.Members)
	assert.Equal(t, len(teamMembersDB), len(responseDTO.Members))
	for i := 0; i < len(teamMembersDB); i++ {
		assert.Equal(t, teamMembersDB[i].TeamName, responseDTO.TeamName)
		assert.Equal(t, usersDB[i].UserID, responseDTO.Members[i].UserID)
		assert.Equal(t, usersDB[i].Username, responseDTO.Members[i].Username)
		assert.Equal(t, usersDB[i].IsActive, responseDTO.Members[i].IsActive)
	}
}
