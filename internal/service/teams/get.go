package serviceteams

import (
	"context"
	"errors"
	"fmt"

	"github.com/sariya23/manage_pr_service/internal/models/domain"
	"github.com/sariya23/manage_pr_service/internal/outerror"
)

func (s *TeamsService) Get(ctx context.Context, teamName string) ([]domain.User, error) {
	const operationPlace = "service.teams.Get"

	_, err := s.teamRepository.GetTeam(ctx, teamName)
	if err != nil {
		if errors.Is(err, outerror.ErrTeamNotFound) {
			return nil, outerror.ErrTeamNotFound
		}
		return nil, fmt.Errorf("%s:%w", operationPlace, err)
	}

	members, err := s.teamRepository.GetTeamMembers(ctx, teamName)
	if err != nil {
		return nil, fmt.Errorf("%s:%w", operationPlace, err)
	}
	return members, nil
}
