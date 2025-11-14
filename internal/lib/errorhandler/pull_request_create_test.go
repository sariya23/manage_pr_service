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

func TestPullRequestCreate(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name        string
		err         error
		wantStatus  int
		wantResp    api.ErrorResponse
		wantIsError bool
	}{
		{
			name:        "no error",
			err:         nil,
			wantStatus:  http.StatusOK,
			wantResp:    api.ErrorResponse{},
			wantIsError: false,
		},
		{
			name:        "pull request already exists",
			err:         outerror.ErrPullRequestAlreadyExists,
			wantStatus:  http.StatusConflict,
			wantResp:    erresponse.MakePullRequestAlreadyExistsResponse("PR id already exists"),
			wantIsError: true,
		},
		{
			name:        "user not found",
			err:         outerror.ErrUserNotFound,
			wantStatus:  http.StatusBadRequest,
			wantResp:    erresponse.MakeNotFoundResponse("author_id not found"),
			wantIsError: true,
		},
		{
			name:        "user not in any team",
			err:         outerror.ErrUserNotInAnyTeam,
			wantStatus:  http.StatusBadRequest,
			wantResp:    erresponse.MakeNotFoundResponse("author_id not in any team"),
			wantIsError: true,
		},
		{
			name:        "unknown error",
			err:         errors.New("random error"),
			wantStatus:  http.StatusInternalServerError,
			wantResp:    erresponse.MakeInternalResponse("internal server error"),
			wantIsError: true,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			st, resp, ok := PullRequestCreate(c.err)
			assert.Equal(t, c.wantStatus, st)
			assert.Equal(t, c.wantResp, resp)
			assert.Equal(t, c.wantIsError, ok)
		})
	}
}
