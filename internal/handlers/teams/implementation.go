package apiteams

import "log/slog"

type TeamService interface{}

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
