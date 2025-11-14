package errorhandler

import (
	"errors"
	"net/http"
	"testing"

	api "github.com/sariya23/manage_pr_service/internal/generated"
	"github.com/sariya23/manage_pr_service/internal/lib/erresponse"
	"github.com/sariya23/manage_pr_service/internal/outerror"
	"github.com/stretchr/testify/assert"
)

func TestPullRequestReassign(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name            string
		inputErr        error
		expectedStatus  int
		expectedResp    api.ErrorResponse
		expectedIsError bool
	}{
		{
			name:            "no error",
			inputErr:        nil,
			expectedStatus:  http.StatusOK,
			expectedResp:    api.ErrorResponse{},
			expectedIsError: false,
		},
		{
			name:            "pull request not found",
			inputErr:        outerror.ErrPullRequestNotFound,
			expectedStatus:  http.StatusNotFound,
			expectedResp:    erresponse.MakeNotFoundResponse("pull request not found"),
			expectedIsError: true,
		},
		{
			name:            "user not found",
			inputErr:        outerror.ErrUserNotFound,
			expectedStatus:  http.StatusNotFound,
			expectedResp:    erresponse.MakeNotFoundResponse("user not found"),
			expectedIsError: true,
		},
		{
			name:            "pull request merged",
			inputErr:        outerror.ErrPullRequestMerged,
			expectedStatus:  http.StatusConflict,
			expectedResp:    erresponse.MakePullRequestMergedResponse("cannot reassign on merged PR"),
			expectedIsError: true,
		},
		{
			name:            "user is not reviewer",
			inputErr:        outerror.ErrUserIsNotReviewer,
			expectedStatus:  http.StatusConflict,
			expectedResp:    erresponse.MakePullRequestUserNotReviewerResponse("reviewer is not assigned to this PR"),
			expectedIsError: true,
		},
		{
			name:            "no reviewer candidates",
			inputErr:        outerror.ErrNoReviewerCandidates,
			expectedStatus:  http.StatusConflict,
			expectedResp:    erresponse.MakePullRequestNoCandidateResponse("no active replacement candidate in team"),
			expectedIsError: true,
		},
		{
			name:            "unknown error",
			inputErr:        errors.New("some unknown error"),
			expectedStatus:  http.StatusInternalServerError,
			expectedResp:    erresponse.MakeInternalResponse("internal server error"),
			expectedIsError: true,
		},
		{
			name:            "user not in PR team",
			inputErr:        outerror.ErrUserNotInPullRequestTeam,
			expectedStatus:  http.StatusBadRequest,
			expectedResp:    erresponse.MakeInvalidResponse("user not in PR team"),
			expectedIsError: true,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			st, resp, isError := PullRequestReassign(c.inputErr)
			assert.Equal(t, c.expectedStatus, st)
			assert.Equal(t, c.expectedResp, resp)
			assert.Equal(t, c.expectedIsError, isError)
		})

	}
}
