//go:build integrations

package users

import "github.com/sariya23/manage_pr_service/tests/clients/postgresql"

var (
	dbT    *postgresql.TestDB
	tables = []string{"\"user\"", "pull_request", "team_member"}
)

func init() {
	dbT = postgresql.NewTestDB()
}
