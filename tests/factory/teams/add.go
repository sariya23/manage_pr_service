package teams

import (
	"github.com/brianvoe/gofakeit/v7"
	"github.com/sariya23/manage_pr_service/tests/helpers/random"
)

type AddTeamRequestMemberDT struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	IsActive bool   `json:"is_active"`
}

func RandomInitAddTeamRequestMemberDT(userID string, username string, isActive *bool) AddTeamRequestMemberDT {
	if userID == "" {
		userID = gofakeit.LetterN(uint(random.RandInt(1, 8)))
	}
	if username == "" {
		username = gofakeit.Username()
	}
	var active bool
	if isActive == nil {
		active = gofakeit.Bool()
	} else {
		active = *isActive
	}

	return AddTeamRequestMemberDT{
		UserID:   userID,
		Username: username,
		IsActive: active,
	}
}

type AddTeamRequest struct {
	TeamName string                   `json:"team_name"`
	Members  []AddTeamRequestMemberDT `json:"members"`
}

func RandomInitAddTeamRequest(teamName string, members []AddTeamRequestMemberDT) AddTeamRequest {
	if teamName == "" {
		teamName = gofakeit.Name()
	}
	membersN := random.RandInt(1, 3)
	resMembers := make([]AddTeamRequestMemberDT, 0, membersN)
	if members == nil {
		for range membersN {
			resMembers = append(resMembers, RandomInitAddTeamRequestMemberDT("", "", nil))
		}
	} else {
		resMembers = members
	}
	return AddTeamRequest{
		TeamName: teamName,
		Members:  resMembers,
	}
}
