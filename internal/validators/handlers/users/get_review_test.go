package validators

import (
	"testing"

	api "github.com/sariya23/manage_pr_service/internal/generated"
	"github.com/stretchr/testify/assert"
)

func TestValidateGetUserReviewRequest(t *testing.T) {
	cases := []struct {
		name          string
		request       api.GetUsersGetReviewRequestObject
		expectedValid bool
		expectedMsg   string
	}{
		{
			name:          "Empty user id",
			request:       api.GetUsersGetReviewRequestObject{Params: api.GetUsersGetReviewParams{UserId: ""}},
			expectedValid: false,
			expectedMsg:   "user_id is required",
		},
		{
			name:          "User id not numeric",
			request:       api.GetUsersGetReviewRequestObject{Params: api.GetUsersGetReviewParams{UserId: "ABOBA"}},
			expectedValid: false,
			expectedMsg:   "user_id must be numeric",
		},
		{
			name:          "User id is negative",
			request:       api.GetUsersGetReviewRequestObject{Params: api.GetUsersGetReviewParams{UserId: "-123"}},
			expectedValid: false,
			expectedMsg:   "user_id must be positive",
		},
		{
			name:          "OK",
			request:       api.GetUsersGetReviewRequestObject{Params: api.GetUsersGetReviewParams{UserId: "22"}},
			expectedValid: true,
			expectedMsg:   "OK",
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			msg, valid := ValidateGetUserReviewRequest(c.request)
			assert.Equal(t, c.expectedValid, valid)
			assert.Equal(t, c.expectedMsg, msg)
		})
	}
}
