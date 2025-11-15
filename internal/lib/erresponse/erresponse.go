package erresponse

import (
	api "github.com/sariya23/manage_pr_service/internal/generated"
	"github.com/sariya23/manage_pr_service/internal/handlers"
)

func MakeInvalidResponse(msg string) api.ErrorResponse {
	errorResp := api.ErrorResponse{}
	errorResp.Error.Code = handlers.INVALIDREQUEST
	errorResp.Error.Message = msg
	return errorResp
}

func MakeNotFoundResponse(msg string) api.ErrorResponse {
	errorResp := api.ErrorResponse{}
	errorResp.Error.Code = handlers.NOTFOUND
	errorResp.Error.Message = msg
	return errorResp
}

func MakeTeamAlreadyExistsResponse(msg string) api.ErrorResponse {
	errorResp := api.ErrorResponse{}
	errorResp.Error.Code = handlers.TEAMALREADYEXISTS
	errorResp.Error.Message = msg
	return errorResp
}

func MakeInternalResponse(msg string) api.ErrorResponse {
	errorResp := api.ErrorResponse{}
	errorResp.Error.Code = handlers.INTERNAL
	errorResp.Error.Message = msg
	return errorResp
}

func MakePullRequestAlreadyExistsResponse(msg string) api.ErrorResponse {
	errorResp := api.ErrorResponse{}
	errorResp.Error.Code = handlers.PULLREQUESTALREADYEXISTS
	errorResp.Error.Message = msg
	return errorResp
}

func MakePullRequestMergedResponse(msg string) api.ErrorResponse {
	errorResp := api.ErrorResponse{}
	errorResp.Error.Code = handlers.PULLREQUESTMERGED
	errorResp.Error.Message = msg
	return errorResp
}

func MakePullRequestUserNotReviewerResponse(msg string) api.ErrorResponse {
	errorResp := api.ErrorResponse{}
	errorResp.Error.Code = handlers.USERNOTREVIEWER
	errorResp.Error.Message = msg
	return errorResp
}

func MakePullRequestNoCandidateResponse(msg string) api.ErrorResponse {
	errorResp := api.ErrorResponse{}
	errorResp.Error.Code = handlers.NOCANDIDATE
	errorResp.Error.Message = msg
	return errorResp
}
