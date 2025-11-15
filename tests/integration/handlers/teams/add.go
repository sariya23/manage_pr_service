package teams

import (
	"context"
	"testing"

	httpcleint "github.com/sariya23/manage_pr_service/tests/clients/http"
)

func TestAddTeam_NewTeamNewUsers(t *testing.T) {
	ctx := context.Background()
	httpClient := httpcleint.NewHTTPClient()
	dbT.SetUp(ctx, t, tables...)
	defer dbT.TearDown(t)

}
