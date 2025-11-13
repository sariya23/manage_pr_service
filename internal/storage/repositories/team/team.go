package repo_team

import (
	"log/slog"

	"github.com/sariya23/manage_pr_service/internal/storage/database"
)

const (
	TeamUserIDField   = "user_id"
	TeamTeamNameField = "team_name"
	TeamTableName     = "teaam"
)

type TeamRepository struct {
	conn *database.Database
	log  *slog.Logger
}

func NewTeamRepository(conn *database.Database, log *slog.Logger) *TeamRepository {
	return &TeamRepository{
		conn: conn,
		log:  log,
	}
}
