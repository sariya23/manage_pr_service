package validators

import (
	"strconv"

	api "github.com/sariya23/manage_pr_service/internal/generated"
)

func ValidateGetUserReviewRequest(request api.GetUsersGetReviewRequestObject) (string, bool) {
	if request.Params.UserId == "" {
		return "missing query param 'user_id'", false
	}

	userID, err := strconv.Atoi(request.Params.UserId)
	if err != nil {
		return "user_id must be numeric", false
	}
	if userID < 0 {
		return "user_id must be positive", false
	}
	return "OK", true
}
