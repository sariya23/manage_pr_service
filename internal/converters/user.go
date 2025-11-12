package converters

import (
	"strconv"

	api "github.com/sariya23/manage_pr_service/internal/generated"
	"github.com/sariya23/manage_pr_service/internal/models"
)

func DomainUserToIsActiveResponseUser(domainUser models.User) api.PostUsersSetIsActive200JSONResponse {
	var res api.PostUsersSetIsActive200JSONResponse
	res.User.UserId = strconv.Itoa(int(domainUser.UserID))
	res.User.TeamName = domainUser.TeamName
	res.User.IsActive = domainUser.IsActive
	return res
}
