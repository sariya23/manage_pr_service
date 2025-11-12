package validators

import (
	"testing"

	api "github.com/sariya23/manage_pr_service/internal/generated"
	"github.com/stretchr/testify/assert"
)

func TestSetIsActive(t *testing.T) {
	cases := []struct {
		name          string
		request       api.PostUsersSetIsActiveJSONRequestBody
		expectedValid bool
		expectedMsg   string
	}{
		{
			name:          "Empty user id",
			request:       api.PostUsersSetIsActiveJSONRequestBody{IsActive: false, UserId: ""},
			expectedValid: false,
			expectedMsg:   "user_id is required",
		},
		{
			name:          "User id not numeric",
			request:       api.PostUsersSetIsActiveJSONRequestBody{IsActive: false, UserId: "ABOBA"},
			expectedValid: false,
			expectedMsg:   "user_id must be numeric",
		},
		{
			name:          "User id is negative",
			request:       api.PostUsersSetIsActiveJSONRequestBody{IsActive: false, UserId: "-123"},
			expectedValid: false,
			expectedMsg:   "user_id must be positive",
		},
		{
			name:          "OK",
			request:       api.PostUsersSetIsActiveJSONRequestBody{IsActive: false, UserId: "123"},
			expectedValid: true,
			expectedMsg:   "OK",
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			msg, valid := ValidateSetIsActiveUserRequest(c.request)
			assert.Equal(t, c.expectedValid, valid)
			assert.Equal(t, c.expectedMsg, msg)
		})
	}
}
