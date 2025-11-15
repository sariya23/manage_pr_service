//go:build integrations

package teams

import (
	"context"
	"net/http"
	"testing"

	"github.com/sariya23/manage_pr_service/tests/checkers"
	httpcleint "github.com/sariya23/manage_pr_service/tests/clients/http"
	"github.com/sariya23/manage_pr_service/tests/factory"
	"github.com/sariya23/manage_pr_service/tests/helpers/random"
	"github.com/sariya23/manage_pr_service/tests/models"
	"github.com/stretchr/testify/require"
)

// TestAddTeam_NewTeamNewUsers тест ручки /api/team/add
// Создание новой команды и новых юзеров
func TestAddTeam_NewTeamNewUsers(t *testing.T) {
	ctx := context.Background()
	httpClient := httpcleint.NewHTTPClient()
	members := []factory.AddTeamRequestMemberDTO{}
	for range random.RandInt(1, 3) {
		isActive := true
		members = append(members, factory.RandomInitAddTeamRequestMemberDT("", "", &isActive))
	}
	request := factory.RandomInitAddTeamRequest("", members)

	response := httpClient.TeamsAdd(request)
	require.Equal(t, http.StatusOK, response.StatusCode)
	responseDTO := factory.AddTeamRFromHTTPResponseOK(response)
	teamMembersDB := dbT.GetTeamMembersByTeamName(ctx, request.TeamName)
	usersDB := dbT.GetUsersFromDB(ctx, models.TeamMemberUserIDs(teamMembersDB))
	checkers.CheckAddTeamResponse(t, responseDTO, teamMembersDB, usersDB)
}

// TestAddTeam_AddUsersToTeam тест ручки /api/team/add
// Добавление новых юзеров в существующую команду
func TestAddTeam_AddUsersToTeam(t *testing.T) {
	ctx := context.Background()
	dbT.SetUp(ctx, t, tables...)
	httpClient := httpcleint.NewHTTPClient()
	// Предварительно создаем команду с юзерами
	nUsers := random.RandInt(1, 3)
	membersInit := make([]factory.AddTeamRequestMemberDTO, 0, nUsers)
	for range nUsers {
		isActive := true
		membersInit = append(membersInit, factory.RandomInitAddTeamRequestMemberDT("", "", &isActive))
	}
	requestCreateTeam := factory.RandomInitAddTeamRequest("", membersInit)
	responseCreateTeam := httpClient.TeamsAdd(requestCreateTeam)
	require.Equal(t, http.StatusOK, responseCreateTeam.StatusCode)

	nUsers = random.RandInt(1, 3)
	newMembers := make([]factory.AddTeamRequestMemberDTO, 0, nUsers)
	for range nUsers {
		isActive := true
		membersInit = append(membersInit, factory.RandomInitAddTeamRequestMemberDT("", "", &isActive))
	}

	requestAddTeam := factory.RandomInitAddTeamRequest(requestCreateTeam.TeamName, newMembers)
	responseAddTeam := httpClient.TeamsAdd(requestAddTeam)
	require.Equal(t, http.StatusOK, responseAddTeam.StatusCode)
	responseDTO := factory.AddTeamRFromHTTPResponseOK(responseAddTeam)
	teamMembersDB := dbT.GetTeamMembersByTeamName(ctx, requestAddTeam.TeamName)
	usersDB := dbT.GetUsersFromDB(ctx, models.TeamMemberUserIDs(teamMembersDB))
	checkers.CheckAddTeamResponse(t, responseDTO, teamMembersDB, usersDB)
}

// TestAddTeam_InActiveUsers тест ручки /api/team/add
// Если хотя бы один пользователь неактивен, то возвращается ошибка
func TestAddTeam_InActiveUsers(t *testing.T) {
	ctx := context.Background()
	dbT.SetUp(ctx, t, tables...)
	httpClient := httpcleint.NewHTTPClient()
	nUsers := random.RandInt(1, 3)
	membersInit := make([]factory.AddTeamRequestMemberDTO, 0, nUsers)
	for range nUsers {
		isActive := true
		membersInit = append(membersInit, factory.RandomInitAddTeamRequestMemberDT("", "", &isActive))
	}
	diactive := false
	membersInit = append(membersInit, factory.RandomInitAddTeamRequestMemberDT("", "", &diactive))
	request := factory.RandomInitAddTeamRequest("", membersInit)
	response := httpClient.TeamsAdd(request)
	require.Equal(t, http.StatusBadRequest, response.StatusCode)
}

// // TestAddTeam_InActiveUsers тест ручки /api/team/add
// // При попытке добавить пользователя из другой команды возвращается ошибка
func TestAddTeam_UserFromAnotherTeam(t *testing.T) {
	ctx := context.Background()
	dbT.SetUp(ctx, t, tables...)
	httpClient := httpcleint.NewHTTPClient()
	nUsers := random.RandInt(1, 3)
	membersInit := make([]factory.AddTeamRequestMemberDTO, 0, nUsers)
	for range nUsers {
		isActive := true
		membersInit = append(membersInit, factory.RandomInitAddTeamRequestMemberDT("", "", &isActive))
	}
	requestCreateTeam := factory.RandomInitAddTeamRequest("", membersInit)
	responseCreateTeam := httpClient.TeamsAdd(requestCreateTeam)
	require.Equal(t, http.StatusOK, responseCreateTeam.StatusCode)
	user := random.Choice(membersInit)
	nUsers = random.RandInt(1, 3)
	membersAdd := make([]factory.AddTeamRequestMemberDTO, 0, nUsers+1)
	membersAdd = append(membersAdd, user)
	for range nUsers {
		isActive := true
		membersAdd = append(membersAdd, factory.RandomInitAddTeamRequestMemberDT("", "", &isActive))
	}
	request := factory.RandomInitAddTeamRequest("", membersAdd)
	response := httpClient.TeamsAdd(request)
	require.Equal(t, http.StatusBadRequest, response.StatusCode)
}

// TestAddTeam_EmptyTeamName тест ручки /api/team/add
// В запросе не передано название команды
func TestAddTeam_EmptyTeamName(t *testing.T) {
	httpClient := httpcleint.NewHTTPClient()
	request := factory.RandomInitAddTeamRequest("", nil)
	request.TeamName = ""
	response := httpClient.TeamsAdd(request)
	require.Equal(t, http.StatusBadRequest, response.StatusCode)
}
