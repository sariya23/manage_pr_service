package validators

import (
	api "github.com/sariya23/manage_pr_service/internal/generated"
)

func ValidateSetIsActiveUserRequest(request api.PostUsersSetIsActiveJSONRequestBody) (string, bool) {
	return ValidateUserID(request.UserId)
}
