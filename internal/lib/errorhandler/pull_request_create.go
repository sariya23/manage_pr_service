package errorhandler

import (
	"errors"
	"net/http"

	api "github.com/sariya23/manage_pr_service/internal/generated"
	"github.com/sariya23/manage_pr_service/internal/lib/erresponse"
	"github.com/sariya23/manage_pr_service/internal/outerror"
)

func PullRequestCreate(err error) (status int, resp api.ErrorResponse, isError bool) {
	if err == nil {
		return http.StatusOK, api.ErrorResponse{}, false
	}

	if errors.Is(err, outerror.ErrPullRequestAlreadyExists) {
		return http.StatusConflict, erresponse.MakePullRequestAlreadyExistsResponse("PR id already exists"), true
	} else if errors.Is(err, outerror.ErrUserNotFound) {
		return http.StatusBadRequest, erresponse.MakeNotFoundResponse("author_id not found"), true
	} else if errors.Is(err, outerror.ErrUserNotInAnyTeam) {
		return http.StatusBadRequest, erresponse.MakeNotFoundResponse("author_id not in any team"), true
	}
	return http.StatusInternalServerError, erresponse.MakeInternalResponse("internal server error"), true

}
