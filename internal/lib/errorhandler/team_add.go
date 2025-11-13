package errorhandler

import (
	"errors"
	"fmt"
	"net/http"

	api "github.com/sariya23/manage_pr_service/internal/generated"
	"github.com/sariya23/manage_pr_service/internal/lib/erresponse"
	"github.com/sariya23/manage_pr_service/internal/outerror"
)

func TeamAdd(err error, teamName string) (status int, resp api.ErrorResponse, isError bool) {
	if err == nil {
		return http.StatusOK, api.ErrorResponse{}, false
	}

	if errors.Is(err, outerror.ErrTeamAlreadyExists) {
		return http.StatusBadRequest, erresponse.MakeTeamAlreadyExistsResponse(fmt.Sprintf("%s already exists", teamName)), true
	} else if errors.Is(err, outerror.ErrUserAlreadyInTeam) {
		return http.StatusBadRequest, erresponse.MakeInvalidResponse(fmt.Sprintf("one of users already in team %s", teamName)), true
	} else if errors.Is(err, outerror.ErrInactiveUser) {
		return http.StatusBadRequest, erresponse.MakeInvalidResponse("one of users is inactive"), true
	}
	return http.StatusInternalServerError, erresponse.MakeInternalResponse("internal server error"), true
}
