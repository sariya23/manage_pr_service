package converters

import (
	api "github.com/sariya23/manage_pr_service/internal/generated"
	"github.com/sariya23/manage_pr_service/internal/models/domain"
)

func DomainUserToIsActiveResponseUser(domainUser domain.User, teamName string) api.User {
	var res api.User
	res.UserId = domainUser.UserID
	res.TeamName = teamName
	res.Username = domainUser.Username
	res.IsActive = domainUser.IsActive
	return res
}

func MultiAddTeamUserToDomainUser(dto []api.TeamMember) []domain.User {
	res := make([]domain.User, 0, len(dto))

	for _, member := range dto {
		res = append(res, domain.User{
			UserID:   member.UserId,
			Username: member.Username,
			IsActive: member.IsActive,
		})
	}
	return res
}

func MultiDomainUserToAddTeamResponse(domainUser []domain.User) []api.TeamMember {
	res := make([]api.TeamMember, 0, len(domainUser))
	for _, user := range domainUser {
		res = append(res, api.TeamMember{
			UserId:   user.UserID,
			Username: user.Username,
			IsActive: user.IsActive,
		})
	}
	return res
}

func MultiDomainUserToGetTeamResponse(domainUser []domain.User) []api.TeamMember {
	res := make([]api.TeamMember, 0, len(domainUser))
	for _, user := range domainUser {
		res = append(res, api.TeamMember{
			UserId:   user.UserID,
			Username: user.Username,
			IsActive: user.IsActive,
		})
	}
	return res
}
