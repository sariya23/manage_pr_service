package serviceteams

import (
	"context"

	"github.com/sariya23/manage_pr_service/internal/models/domain"
)

func (s *TeamsService) Get(ctx context.Context, teamName string) ([]domain.User, error) {
	//const operationPlace = "service.teams.Get"
	//log := s.log.With("operationPlace", operationPlace)
	panic("implement me")
}
