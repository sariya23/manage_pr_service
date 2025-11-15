package validators

func ValidateTeamGet(teamName string) (string, bool) {
	if teamName == "" {
		return "team_name is required", false
	}

	return "OK", true
}
