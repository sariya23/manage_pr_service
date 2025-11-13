package repo_team

import (
	"context"

	"github.com/sariya23/manage_pr_service/internal/models/domain"
)

func (r *TeamRepository) UpsertTeam(ctx context.Context, teamName string, users []domain.User) error {
	const operationPlace = "storage.repositories.team.UpsertTeam"
	panic("")
}
