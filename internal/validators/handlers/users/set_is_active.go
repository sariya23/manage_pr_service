package validators

import (
	"strconv"

	api "github.com/sariya23/manage_pr_service/internal/generated"
)

func ValidateSetIsActiveUserRequest(request api.PostUsersSetIsActiveJSONRequestBody) (string, bool) {
	if request.UserId == "" {
		return "user_id is required", false
	}

	userID, err := strconv.Atoi(request.UserId)
	if err != nil {
		return "user_id must be numeric", false
	}
	if userID < 0 {
		return "user_id must be positive", false
	}
	return "OK", true
}
