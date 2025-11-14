package errorhandler

import (
	"errors"
	"net/http"

	api "github.com/sariya23/manage_pr_service/internal/generated"
	"github.com/sariya23/manage_pr_service/internal/lib/erresponse"
	"github.com/sariya23/manage_pr_service/internal/outerror"
)

func PullRequestReassign(err error) (status int, resp api.ErrorResponse, isError bool) {
	if err == nil {
		return http.StatusOK, api.ErrorResponse{}, false
	}
	if errors.Is(err, outerror.ErrPullRequestNotFound) {
		return http.StatusNotFound, erresponse.MakeNotFoundResponse("pull request not found"), true
	} else if errors.Is(err, outerror.ErrUserNotFound) {
		return http.StatusNotFound, erresponse.MakeNotFoundResponse("user not found"), true
	} else if errors.Is(err, outerror.ErrPullRequestMerged) {
		return http.StatusConflict, erresponse.MakePullRequestMergedResponse("cannot reassign on merged PR"), true
	} else if errors.Is(err, outerror.ErrUserIsNotReviewer) {
		return http.StatusConflict, erresponse.MakePullRequestUserNotReviewerResponse("reviewer is not assigned to this PR"), true
	} else if errors.Is(err, outerror.ErrNoReviewerCandidates) {
		return http.StatusConflict, erresponse.MakePullRequestNoCandidateResponse("no active replacement candidate in team"), true
	}

	return http.StatusInternalServerError, erresponse.MakeInternalResponse("internal server error"), true
}
