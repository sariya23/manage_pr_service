//go:build integrations

package users

import (
	"context"
	"net/http"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	httpcleint "github.com/sariya23/manage_pr_service/tests/clients/http"
	factory_teams "github.com/sariya23/manage_pr_service/tests/factory/teams"
	factory_users "github.com/sariya23/manage_pr_service/tests/factory/users"
	"github.com/sariya23/manage_pr_service/tests/helpers/random"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestSetIsActive тест ручки /api/users/setIsActive
// Успешная деактивация пользователя
func TestSetIsActive(t *testing.T) {
	ctx := context.Background()
	httpClient := httpcleint.NewHTTPClient()
	nUsers := random.RandInt(1, 3)
	members := make([]factory_teams.AddTeamRequestMemberDTO, 0, nUsers)
	for range nUsers {
		isActive := true
		members = append(members, factory_teams.RandomInitAddTeamRequestMemberDT("", "", &isActive))
	}
	requestCreate := factory_teams.RandomInitAddTeamRequest("", members)
	responseCreate := httpClient.TeamsAdd(requestCreate)
	require.Equal(t, http.StatusOK, responseCreate.StatusCode)

	user := random.Choice(members)
	request := factory_users.SetIsActiveRequest{UserID: user.UserID, IsActive: false}
	response := httpClient.UsersSetIsActive(request)
	require.Equal(t, http.StatusOK, response.StatusCode)
	responseDTO := factory_users.SetIsActiveFromHTTPResponseOK(response)
	userDB := dbT.GetUsersFromDB(ctx, []string{user.UserID})[0]

	assert.Equal(t, userDB.IsActive, request.IsActive)

	assert.Equal(t, userDB.UserID, responseDTO.User.UserID)
	assert.Equal(t, userDB.Username, responseDTO.User.Username)
	assert.Equal(t, requestCreate.TeamName, responseDTO.User.TeamName)
	assert.Equal(t, userDB.IsActive, responseDTO.User.IsActive)
}

// TestSetIsActive_NonexistentUser тест ручки /api/users/setIsActive
// Ошибка, при попытке обновить несущесвующего пользователя
func TestSetIsActive_NonexistentUser(t *testing.T) {
	httpClient := httpcleint.NewHTTPClient()
	request := factory_users.SetIsActiveRequest{UserID: gofakeit.LetterN(8), IsActive: false}
	response := httpClient.UsersSetIsActive(request)
	require.Equal(t, http.StatusNotFound, response.StatusCode)
}

// TestSetIsActive_EmptyUserID тест ручки /api/users/setIsActive
// Пустой айдишник пользователя
func TestSetIsActive_EmptyUserID(t *testing.T) {
	httpClient := httpcleint.NewHTTPClient()
	request := factory_users.SetIsActiveRequest{UserID: "", IsActive: false}
	response := httpClient.UsersSetIsActive(request)
	require.Equal(t, http.StatusBadRequest, response.StatusCode)
}
