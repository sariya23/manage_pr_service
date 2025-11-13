package validators

func ValidateUserID(userIDBody string) (string, bool) {
	if userIDBody == "" {
		return "user_id is required", false
	}

	return "OK", true
}
