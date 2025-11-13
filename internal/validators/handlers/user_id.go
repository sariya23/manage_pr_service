package validators

import (
	"strconv"
)

func ValidateUserID(userIDBody string) (string, bool) {
	if userIDBody == "" {
		return "user_id is required", false
	}

	userID, err := strconv.Atoi(userIDBody)
	if err != nil {
		return "user_id must be numeric", false
	}
	if userID < 0 {
		return "user_id must be positive", false
	}

	return "OK", true
}
