package erresponse

import (
	api "github.com/sariya23/manage_pr_service/internal/generated"
	"github.com/sariya23/manage_pr_service/internal/handlers/codes"
)

func MakeInvalidResponse(msg string) api.ErrorResponse {
	errorResp := api.ErrorResponse{}
	errorResp.Error.Code = codes.INVALIDREQUEST
	errorResp.Error.Message = msg
	return errorResp
}

func MakeNotFoundResponse(msg string) api.ErrorResponse {
	errorResp := api.ErrorResponse{}
	errorResp.Error.Code = codes.NOTFOUND
	errorResp.Error.Message = msg
	return errorResp
}

func MakeTeamAlreadyExistsResponse(msg string) api.ErrorResponse {
	errorResp := api.ErrorResponse{}
	errorResp.Error.Code = codes.TEAMALREADYEXISTS
	errorResp.Error.Message = msg
	return errorResp
}

func MakeInternalResponse(msg string) api.ErrorResponse {
	errorResp := api.ErrorResponse{}
	errorResp.Error.Code = codes.INTERNAL
	errorResp.Error.Message = msg
	return errorResp
}

func MakePullRequestAlreadyExistsResponse(msg string) api.ErrorResponse {
	errorResp := api.ErrorResponse{}
	errorResp.Error.Code = codes.PULLREQUESTALREADYEXISTS
	errorResp.Error.Message = msg
	return errorResp
}

func MakePullRequestMergedResponse(msg string) api.ErrorResponse {
	errorResp := api.ErrorResponse{}
	errorResp.Error.Code = codes.PULLREQUESTMERGED
	errorResp.Error.Message = msg
	return errorResp
}

func MakePullRequestUserNotReviewerResponse(msg string) api.ErrorResponse {
	errorResp := api.ErrorResponse{}
	errorResp.Error.Code = codes.USERNOTREVIEWER
	errorResp.Error.Message = msg
	return errorResp
}

func MakePullRequestNoCandidateResponse(msg string) api.ErrorResponse {
	errorResp := api.ErrorResponse{}
	errorResp.Error.Code = codes.NOCANDIDATE
	errorResp.Error.Message = msg
	return errorResp
}
