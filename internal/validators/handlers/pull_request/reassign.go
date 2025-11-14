package pull_request_validators

import api "github.com/sariya23/manage_pr_service/internal/generated"

func ValidatePullRequestReassignRequest(request api.PostPullRequestReassignJSONRequestBody) (string, bool) {
	if request.OldUserId == "" {
		return "old_reviewer_id is required", false
	}
	if request.PullRequestId == "" {
		return "pull_request_id is required", false
	}
	return "OK", true
}
