package teams

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"sort"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/sariya23/manage_pr_service/tests/helpers/random"
)

type AddTeamRequestMemberDTO struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	IsActive bool   `json:"is_active"`
}

func RandomInitAddTeamRequestMemberDT(userID string, username string, isActive *bool) AddTeamRequestMemberDTO {
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

	return AddTeamRequestMemberDTO{
		UserID:   userID,
		Username: username,
		IsActive: active,
	}
}

type AddTeamRequest struct {
	TeamName string                    `json:"team_name"`
	Members  []AddTeamRequestMemberDTO `json:"members"`
}

func (r *AddTeamRequest) ToJson() io.Reader {
	const operationPlace = "factory.teams.add.AddTeamRequest.ToJson"
	body, err := json.Marshal(r)
	if err != nil {
		panic(err.Error() + " " + operationPlace)
	}
	return bytes.NewBuffer(body)
}

func RandomInitAddTeamRequest(teamName string, members []AddTeamRequestMemberDTO) AddTeamRequest {
	if teamName == "" {
		teamName = gofakeit.Name()
	}
	membersN := random.RandInt(1, 3)
	resMembers := make([]AddTeamRequestMemberDTO, 0, membersN)
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

type AddTeamResponseMemberDTO struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	IsActive bool   `json:"is_active"`
}

func AddTeamResponseMemberSortByUserID(teamMembers []AddTeamResponseMemberDTO) {
	sort.Slice(teamMembers, func(i, j int) bool {
		return teamMembers[i].UserID < teamMembers[j].UserID
	})
}

type AddTeamResponseTeamDTO struct {
	TeamName string                     `json:"team_name"`
	Members  []AddTeamResponseMemberDTO `json:"members"`
}

type AddTeamResponse struct {
	Team AddTeamResponseTeamDTO `json:"team"`
}

func FromHTTPResponseOK(resp *http.Response) AddTeamResponse {
	const operationPlace = "factory.teams.add.FromHTTPResponse"
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error() + " " + operationPlace)
	}

	var result AddTeamResponse
	if err := json.Unmarshal(body, &result); err != nil {
		panic(err.Error() + " " + operationPlace)
	}
	return result
}
