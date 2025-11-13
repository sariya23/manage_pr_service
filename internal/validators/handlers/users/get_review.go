package usersvalidators

import (
	api "github.com/sariya23/manage_pr_service/internal/generated"
	validators "github.com/sariya23/manage_pr_service/internal/validators/handlers"
)

func ValidateGetUserReviewRequest(request api.GetUsersGetReviewRequestObject) (string, bool) {
	return validators.ValidateUserID(request.Params.UserId)
}
