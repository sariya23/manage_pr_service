package teamsvalidators

import (
	api "github.com/sariya23/manage_pr_service/internal/generated"
	validators "github.com/sariya23/manage_pr_service/internal/validators/handlers"
)

func ValidateTeamAddRequest(request api.PostTeamAddJSONRequestBody) (string, bool) {
	if request.TeamName == "" {
		return "teamname is required", false
	}

	for _, member := range request.Members {
		if member.Username == "" {
			return "username is required", false
		}
		if msg, valid := validators.ValidateUserID(member.UserId); !valid {
			return msg, false
		}
	}
	return "OK", true
}
