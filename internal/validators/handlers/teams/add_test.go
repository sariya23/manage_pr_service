package teamsvalidators

import (
	"testing"

	api "github.com/sariya23/manage_pr_service/internal/generated"
	"github.com/stretchr/testify/assert"
)

func TestValidateTeamAddRequest(t *testing.T) {
	cases := []struct {
		name          string
		request       api.PostTeamAddJSONRequestBody
		expectedMsg   string
		expectedValid bool
	}{
		{
			name: "valid",
			request: api.PostTeamAddJSONRequestBody{
				TeamName: "test",
				Members:  []api.TeamMember{{IsActive: true, Username: "test", UserId: "123"}},
			},
			expectedValid: true,
			expectedMsg:   "OK",
		},
		{
			name: "invalid team name",
			request: api.PostTeamAddJSONRequestBody{
				TeamName: "",
				Members:  []api.TeamMember{{IsActive: true, Username: "test", UserId: "123"}},
			},
			expectedValid: false,
			expectedMsg:   "teamname is required",
		},
		{
			name: "invalid username",
			request: api.PostTeamAddJSONRequestBody{
				TeamName: "Test",
				Members: []api.TeamMember{
					{IsActive: true, Username: "test", UserId: "123"},
					{IsActive: true, Username: "", UserId: "123"},
				},
			},
			expectedValid: false,
			expectedMsg:   "username is required",
		},
		{
			name: "no user_id",
			request: api.PostTeamAddJSONRequestBody{
				TeamName: "Test",
				Members: []api.TeamMember{
					{IsActive: true, Username: "test", UserId: ""},
					{IsActive: true, Username: "test", UserId: "123"},
				},
			},
			expectedValid: false,
			expectedMsg:   "user_id is required",
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			msg, valid := ValidateTeamAddRequest(tc.request)
			assert.Equal(t, tc.expectedValid, valid)
			assert.Equal(t, tc.expectedMsg, msg)
		})
	}
}
