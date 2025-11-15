//go:build integrations

package postgresql

import (
	"context"
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
	"testing"

	"github.com/sariya23/manage_pr_service/internal/config"
	"github.com/sariya23/manage_pr_service/internal/storage/database"
	"github.com/sariya23/manage_pr_service/tests/factory"
)

type TestDB struct {
	DB *database.Database
}

func NewTestDB() *TestDB {
	const operationPlace = "clients.postgresql.NewTestDB"
	cfg := config.MustLoadByPath(filepath.Join("..", "..", "..", "..", "config", "test.env"))
	DB, err := database.NewConnection(
		context.Background(),
		database.GenerateDBUrl(
			cfg.PostgresUsername,
			cfg.PostgresPassword,
			cfg.PostgresOuterHost,
			strconv.Itoa(cfg.PostgresPort),
			cfg.PostgresDB,
			cfg.SSLMode,
		),
	)
	if err != nil {
		panic(err.Error() + " " + operationPlace)
	}
	return &TestDB{DB: DB}
}

func (d *TestDB) TearDown(t *testing.T) {
	t.Helper()
}

func (d *TestDB) Truncate(ctx context.Context, tables ...string) {
	q := fmt.Sprintf("truncate %s cascade", strings.Join(tables, ","))
	if _, err := d.DB.GetPool().Exec(ctx, q); err != nil {
		panic(err)
	}
}

// TEAMS
func (d *TestDB) GetTeamMembersByTeamName(ctx context.Context, teamName string) []factory.TeamMember {
	const operationPlace = "clients.postgresql.NewTestDB"
	getTeamMembersSQL := `select * from team_member where team_name=$1`

	rows, err := d.DB.GetPool().Query(ctx, getTeamMembersSQL, teamName)
	if err != nil {
		panic(err.Error() + " " + operationPlace)
	}
	defer rows.Close()
	var teamMembers []factory.TeamMember

	for rows.Next() {
		var teamMember factory.TeamMember
		err = rows.Scan(
			&teamMember.TeamName,
			&teamMember.UserID)
		if err != nil {
			panic(err.Error() + " " + operationPlace)
		}
		if rows.Err() != nil {
			panic(rows.Err().Error() + " " + operationPlace)
		}
		teamMembers = append(teamMembers, teamMember)
	}
	return teamMembers
}

// USER

func (d *TestDB) GetUsersFromDB(ctx context.Context, userIDs []string) []factory.User {
	const operationPlace = "clients.postgresql.GetUsersFromDB"

	getUsersSQL := `select * from "user" where user_id=any($1)`

	rows, err := d.DB.GetPool().Query(ctx, getUsersSQL, userIDs)
	if err != nil {
		panic(err.Error() + " " + operationPlace)
	}
	defer rows.Close()
	var users []factory.User
	for rows.Next() {
		var user factory.User
		err = rows.Scan(
			&user.UserID,
			&user.Username,
			&user.IsActive,
			&user.CreatedAt,
			&user.UpdatedAt)
		if err != nil {
			panic(err.Error() + " " + operationPlace)
		}
		if rows.Err() != nil {
			panic(rows.Err().Error() + " " + operationPlace)
		}
		users = append(users, user)
	}
	return users
}
