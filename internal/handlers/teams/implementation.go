package apiteams

import (
	"context"
	"log/slog"

	"github.com/sariya23/manage_pr_service/internal/models/domain"
	"github.com/sariya23/manage_pr_service/internal/models/dto"
)

type TeamService interface {
	Add(ctx context.Context, teamName string, members []dto.TeamMember) ([]domain.User, error)
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
