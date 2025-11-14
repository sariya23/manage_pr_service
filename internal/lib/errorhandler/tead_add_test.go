package errorhandler

import (
	"errors"
	"fmt"
	"net/http"
	"testing"

	api "github.com/sariya23/manage_pr_service/internal/generated"
	"github.com/sariya23/manage_pr_service/internal/lib/erresponse"
	"github.com/sariya23/manage_pr_service/internal/outerror"
	"github.com/stretchr/testify/assert"
)

func TestTeamAdd(t *testing.T) {
	t.Parallel()
	someErr := errors.New("some error")
	cases := []struct {
		name            string
		teamName        string
		err             error
		expectedStatus  int
		expectedResp    api.ErrorResponse
		expectedIsError bool
	}{
		{
			name:            "no error",
			teamName:        "test",
			err:             nil,
			expectedStatus:  http.StatusOK,
			expectedResp:    api.ErrorResponse{},
			expectedIsError: false,
		},
		{
			name:            "team already exists",
			teamName:        "test",
			err:             outerror.ErrTeamAlreadyExists,
			expectedStatus:  http.StatusBadRequest,
			expectedResp:    erresponse.MakeTeamAlreadyExistsResponse(fmt.Sprintf("%s already exists", "test")),
			expectedIsError: true,
		},
		{
			name:            "user in another team",
			teamName:        "test",
			err:             outerror.ErrUserAlreadyInTeam,
			expectedStatus:  http.StatusBadRequest,
			expectedResp:    erresponse.MakeInvalidResponse(fmt.Sprintf("one of users already in team %s", "test")),
			expectedIsError: true,
		},
		{
			name:            "inactive user",
			teamName:        "test",
			err:             outerror.ErrInactiveUser,
			expectedStatus:  http.StatusBadRequest,
			expectedResp:    erresponse.MakeInvalidResponse("one of users is inactive"),
			expectedIsError: true,
		},
		{
			name:            "internal",
			teamName:        "test",
			err:             someErr,
			expectedStatus:  http.StatusInternalServerError,
			expectedResp:    erresponse.MakeInternalResponse("internal server error"),
			expectedIsError: true,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			st, resp, isError := TeamAdd(tc.err, tc.teamName)
			assert.Equal(t, tc.expectedStatus, st)
			assert.Equal(t, tc.expectedResp, resp)
			assert.Equal(t, tc.expectedIsError, isError)
		})
	}
}
