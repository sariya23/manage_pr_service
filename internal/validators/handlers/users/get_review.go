package usersvalidators

import (
	validators "github.com/sariya23/manage_pr_service/internal/validators/handlers"
)

func ValidateGetUserReviewRequest(userID string) (string, bool) {
	return validators.ValidateUserID(userID)
}
