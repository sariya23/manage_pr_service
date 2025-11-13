package repo_team

import (
	"github.com/sariya23/manage_pr_service/internal/storage/database"
)

const (
	TeamTableName                = "team"
	TeamMemberTableName          = "team_member"
	TeamTableTeamNameField       = "team_name"
	TeamMemberTableTeamNameField = "team_name"
	TeamMemberTableUserIDField   = "user_id"
)

type TeamRepository struct {
	conn *database.Database
}

func NewTeamRepository(conn *database.Database) *TeamRepository {
	return &TeamRepository{
		conn: conn,
	}
}
