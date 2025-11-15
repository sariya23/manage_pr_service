package validators

import (
	api "github.com/sariya23/manage_pr_service/internal/generated"
)

func ValidatePullRequestCreateRequest(request api.PostPullRequestCreateJSONRequestBody) (string, bool) {
	if request.PullRequestId == "" {
		return "pull_request_id is required", false
	}
	if request.PullRequestName == "" {
		return "pull_request_name is required", false
	}

	if msg, valid := ValidateUserID(request.AuthorId); !valid {
		return msg, false
	}
	return "OK", true
}
