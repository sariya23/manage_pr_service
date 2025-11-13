package erresponse

import (
	api "github.com/sariya23/manage_pr_service/internal/generated"
	"github.com/sariya23/manage_pr_service/internal/handlers"
)

func MakeInvalidResponse(msg string) api.ErrorResponse {
	errorResp := api.ErrorResponse{}
	errorResp.Error.Code = handlers.INVALID_REQUEST
	errorResp.Error.Message = msg
	return errorResp
}

func MakeNotFoundResponse(msg string) api.ErrorResponse {
	errorResp := api.ErrorResponse{}
	errorResp.Error.Code = handlers.NOT_FOUND
	errorResp.Error.Message = msg
	return errorResp
}

func MakeTeamAlreadyExistsResponse(msg string) api.ErrorResponse {
	errorResp := api.ErrorResponse{}
	errorResp.Error.Code = handlers.TEAM_ALREADY_EXISTS
	errorResp.Error.Message = msg
	return errorResp
}

func MakeInternalResponse(msg string) api.ErrorResponse {
	errorResp := api.ErrorResponse{}
	errorResp.Error.Code = handlers.INTERNAL
	errorResp.Error.Message = msg
	return errorResp
}
