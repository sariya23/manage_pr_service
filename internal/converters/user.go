package converters

import (
	"strconv"

	api "github.com/sariya23/manage_pr_service/internal/generated"
	"github.com/sariya23/manage_pr_service/internal/models"
)

func DomainUserToIsActiveResponseUser(domainUser models.User) api.User {
	var res api.User
	res.UserId = strconv.Itoa(int(domainUser.UserID))
	res.TeamName = domainUser.TeamName
	res.IsActive = domainUser.IsActive
	return res
}
