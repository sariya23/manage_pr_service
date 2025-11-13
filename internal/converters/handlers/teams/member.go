package teamsconverters

import (
	"strconv"

	api "github.com/sariya23/manage_pr_service/internal/generated"
	"github.com/sariya23/manage_pr_service/internal/models/dto"
)

func MultiToDTOMember(members []api.TeamMember) []dto.TeamMember {
	res := make([]dto.TeamMember, 0, len(members))

	for _, member := range members {
		userID, _ := strconv.Atoi(member.UserId)
		res = append(res, dto.TeamMember{
			Username: member.Username,
			IsActive: member.IsActive,
			UserID:   int64(userID),
		})
	}
	return res
}
