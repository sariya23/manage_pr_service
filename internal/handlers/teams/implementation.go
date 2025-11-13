package apiteams

import (
	"context"
	"log/slog"

	"github.com/sariya23/manage_pr_service/internal/models/domain"
)

type TeamService interface {
	Add(ctx context.Context, teamName string, members []domain.User) ([]domain.User, error)
	Get(ctx context.Context, teamName string) ([]domain.User, error)
}

type TeamsImplementation struct {
	logger       *slog.Logger
	teamsService TeamService
}

func NewTeamsImplementation(log *slog.Logger, teamsSrv TeamService) *TeamsImplementation {
	return &TeamsImplementation{
		logger:       log,
		teamsService: teamsSrv,
	}
}
