//go:build integrations

package factory

import (
	"encoding/json"
	"io"
	"net/http"
	"sort"
)

type GetTeamResponseMemberDTO struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	IsActive bool   `json:"is_active"`
}

func GetTeamResponseMemberSortByUserID(teamMembers []GetTeamResponseMemberDTO) {
	sort.Slice(teamMembers, func(i, j int) bool {
		return teamMembers[i].UserID < teamMembers[j].UserID
	})
}

type GetTeamResponse struct {
	TeamName string                     `json:"team_name"`
	Members  []GetTeamResponseMemberDTO `json:"members"`
}

func GetTeamRFromHTTPResponseOK(resp *http.Response) GetTeamResponse {
	const operationPlace = "factory.teams.GetTeamRFromHTTPResponseOK"
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error() + " " + operationPlace)
	}

	var result GetTeamResponse
	if err := json.Unmarshal(body, &result); err != nil {
		panic(err.Error() + " " + operationPlace)
	}
	return result
}
