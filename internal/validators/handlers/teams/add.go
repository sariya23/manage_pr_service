package validators

import (
	api "github.com/sariya23/manage_pr_service/internal/generated"
	validators "github.com/sariya23/manage_pr_service/internal/validators/handlers"
)

func ValidateTeamAddRequest(request api.PostTeamAddJSONRequestBody) (string, bool) {
	if request.TeamName == "" {
		return "teamname is required", false
	}

	if len(request.Members) == 0 {
		return "must have at least one member", false
	}

	for _, member := range request.Members {
		if msg, valid := validators.ValidateUserID(member.UserId); !valid {
			return msg, false
		}
		if member.Username == "" {
			return "username is required", false
		}
	}
	return "OK", true
}
