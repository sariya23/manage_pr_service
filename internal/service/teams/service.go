package serviceteams

import (
	"context"
	"log/slog"

	"github.com/sariya23/manage_pr_service/internal/models/domain"
)

type TeamRepository interface {
	InAnyTeam(ctx context.Context, userID string) (bool, error)
	IsExists(context.Context, string) (bool, error)
	GetTeamUserIDs(ctx context.Context, teamName string) ([]string, error)
	UpsertTeam(ctx context.Context, teamName string, users []domain.User) error
}

type UserRepository interface {
	IsExists(ctx context.Context, userID string) (bool, error)
	CreateUser(ctx context.Context, user domain.User) (domain.User, error)
}

type TeamsService struct {
	log            *slog.Logger
	teamRepository TeamRepository
	userRepository UserRepository
}

func NewTeamsService(log *slog.Logger, teamRepo TeamRepository, userRepo UserRepository) *TeamsService {
	return &TeamsService{
		log:            log,
		teamRepository: teamRepo,
		userRepository: userRepo,
	}
}
