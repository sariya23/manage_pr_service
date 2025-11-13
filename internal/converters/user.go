package converters

import (
	api "github.com/sariya23/manage_pr_service/internal/generated"
	"github.com/sariya23/manage_pr_service/internal/models/domain"
)

func DomainUserToIsActiveResponseUser(domainUser domain.User) api.User {
	var res api.User
	res.UserId = domainUser.UserID
	res.TeamName = domainUser.TeamName
	res.IsActive = domainUser.IsActive
	return res
}
