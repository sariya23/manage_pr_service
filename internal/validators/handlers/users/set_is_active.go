package usersvalidators

import (
	api "github.com/sariya23/manage_pr_service/internal/generated"
	validators "github.com/sariya23/manage_pr_service/internal/validators/handlers"
)

func ValidateSetIsActiveUserRequest(request api.PostUsersSetIsActiveJSONRequestBody) (string, bool) {
	return validators.ValidateUserID(request.UserId)
}
