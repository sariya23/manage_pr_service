package validators

func ValidateGetUserReviewRequest(userID string) (string, bool) {
	return ValidateUserID(userID)
}
