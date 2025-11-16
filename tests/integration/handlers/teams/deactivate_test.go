//go:build integrations

package teams

import (
	"context"
	"net/http"
	"testing"

	httpclient "github.com/sariya23/manage_pr_service/tests/clients/http"
	"github.com/sariya23/manage_pr_service/tests/factory"
	"github.com/sariya23/manage_pr_service/tests/helpers/random"
	"github.com/sariya23/manage_pr_service/tests/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestDeactivate тест на ручку teams/deactivate
// Успешная деактивация части пользователей команды
func TestDeactivate(t *testing.T) {
	ctx := context.Background()
	httpClient := httpclient.NewHTTPClient()
	members := []factory.AddTeamRequestMemberDTO{}
	for range random.RandInt(3, 5) {
		isActive := true
		members = append(members, factory.RandomInitAddTeamRequestMemberDT("", "", &isActive))
	}
	requestCreateTeam := factory.RandomInitAddTeamRequest("", members)
	responseCreateTeam := httpClient.TeamsAdd(requestCreateTeam)
	require.Equal(t, http.StatusOK, responseCreateTeam.StatusCode, responseCreateTeam.Header.Get("requestID"))

	teamMembers := dbT.GetTeamMembersByTeamName(ctx, requestCreateTeam.TeamName)
	teamMembersBefore := random.Sample(teamMembers, 2)
	requestDeactivate := factory.TeamsDeactivateRequest{TeamName: requestCreateTeam.TeamName,
		UserIDs: models.TeamMemberUserIDs(teamMembersBefore),
	}
	responseDeactivate := httpClient.TeamDeactivate(requestDeactivate)
	require.Equal(t, http.StatusOK, responseDeactivate.StatusCode, responseDeactivate.Header.Get("requestID"))

	teamMembersAfter := dbT.GetTeamMembersByTeamName(ctx, requestDeactivate.TeamName)
	usersAfter := dbT.GetUsersFromDB(ctx, models.TeamMemberUserIDs(teamMembersAfter))

	for _, teamMember := range teamMembersBefore {
		user := models.GetByID(usersAfter, teamMember.UserID)
		assert.False(t, user.IsActive, responseDeactivate.Header.Get("requestID"))
	}
}

// TestDeactivate_NonexistentTeam тест ручки /team/deactivate
// Ошибка при попытке деактивировать пользователей в несуществующей команде
func TestDeactivate_NonexistentTeam(t *testing.T) {
	requestDeactivate := factory.TeamsDeactivateRequest{TeamName: "aboba",
		UserIDs: []string{"zxc"},
	}
	httpClient := httpclient.NewHTTPClient()
	responseDeactivate := httpClient.TeamDeactivate(requestDeactivate)
	require.Equal(t, http.StatusNotFound, responseDeactivate.StatusCode, responseDeactivate.Header.Get("requestID"))
}

// TestDeactivate_NonexistentTeam тест ручки /team/deactivate
// Ошибка при попытке деактивировать несуществующего пользователя
func TestDeactivate_NonexistentUser(t *testing.T) {
	httpClient := httpclient.NewHTTPClient()
	members := []factory.AddTeamRequestMemberDTO{}
	for range random.RandInt(1, 2) {
		isActive := true
		members = append(members, factory.RandomInitAddTeamRequestMemberDT("", "", &isActive))
	}
	requestCreateTeam := factory.RandomInitAddTeamRequest("", members)
	responseCreateTeam := httpClient.TeamsAdd(requestCreateTeam)
	require.Equal(t, http.StatusOK, responseCreateTeam.StatusCode, responseCreateTeam.Header.Get("requestID"))
	requestDeactivate := factory.TeamsDeactivateRequest{TeamName: requestCreateTeam.TeamName,
		UserIDs: []string{"zxc"},
	}
	responseDeactivate := httpClient.TeamDeactivate(requestDeactivate)
	require.Equal(t, http.StatusBadRequest, responseDeactivate.StatusCode, responseDeactivate.Header.Get("requestID"))
}

// TestDeactivate_UserNotInThisTeam тест ручки /team/deactivate
// Ошибка при попытке деактивировать пользователя не из команды запроса
func TestDeactivate_UserNotInThisTeam(t *testing.T) {
	ctx := context.Background()
	httpClient := httpclient.NewHTTPClient()
	members1 := []factory.AddTeamRequestMemberDTO{}
	for range random.RandInt(1, 2) {
		isActive := true
		members1 = append(members1, factory.RandomInitAddTeamRequestMemberDT("", "", &isActive))
	}
	requestCreateTeam1 := factory.RandomInitAddTeamRequest("", members1)
	responseCreateTeam1 := httpClient.TeamsAdd(requestCreateTeam1)
	require.Equal(t, http.StatusOK, responseCreateTeam1.StatusCode, responseCreateTeam1.Header.Get("requestID"))

	members2 := []factory.AddTeamRequestMemberDTO{}
	for range random.RandInt(1, 2) {
		isActive := true
		members2 = append(members2, factory.RandomInitAddTeamRequestMemberDT("", "", &isActive))
	}
	requestCreateTeam2 := factory.RandomInitAddTeamRequest("", members2)
	responseCreateTeam2 := httpClient.TeamsAdd(requestCreateTeam2)
	require.Equal(t, http.StatusOK, responseCreateTeam2.StatusCode, responseCreateTeam2.Header.Get("requestID"))

	teamMembersDB := dbT.GetTeamMembersByTeamName(ctx, requestCreateTeam2.TeamName)
	requestDeactivate := factory.TeamsDeactivateRequest{TeamName: requestCreateTeam1.TeamName,
		UserIDs: models.TeamMemberUserIDs(teamMembersDB),
	}
	responseDeactivate := httpClient.TeamDeactivate(requestDeactivate)
	require.Equal(t, http.StatusBadRequest, responseDeactivate.StatusCode, responseDeactivate.Header.Get("requestID"))
}

// TestDeactivate_ValidationError тест на ручку /team/deactivate
// Ошибки валидации
func TestDeactivate_ValidationError(t *testing.T) {
	httpClient := httpclient.NewHTTPClient()
	cases := []struct {
		name string
		req  factory.TeamsDeactivateRequest
	}{
		{
			name: "empty teamName",
			req:  factory.TeamsDeactivateRequest{UserIDs: []string{"asd"}},
		},
		{
			name: "empty userIDs",
			req:  factory.TeamsDeactivateRequest{TeamName: "qwe"},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			response := httpClient.TeamDeactivate(c.req)
			require.Equal(t, http.StatusBadRequest, response.StatusCode, response.Header.Get("requestID"))
		})
	}
}
