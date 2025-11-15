package validators

import (
	api "github.com/sariya23/manage_pr_service/internal/generated"
)

func ValidateTeamAddRequest(request api.PostTeamAddJSONRequestBody) (string, bool) {
	if request.TeamName == "" {
		return "teamname is required", false
	}

	for _, member := range request.Members {
		if member.Username == "" {
			return "username is required", false
		}
		if msg, valid := ValidateUserID(member.UserId); !valid {
			return msg, false
		}
	}
	return "OK", true
}
