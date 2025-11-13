package serviceteams

import (
	"context"
	"log/slog"

	"github.com/sariya23/manage_pr_service/internal/models/domain"
)

type TeamRepository interface {
	GetUserTeam(ctx context.Context, userID string) (string, error)
	GetTeamMemberIDs(ctx context.Context, teamName string) ([]string, error)
	InsertTeam(ctx context.Context, teamName string, users []domain.User) error
}

type UserRepository interface {
	GetUserByID(ctx context.Context, userID string) (*domain.User, error)
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
