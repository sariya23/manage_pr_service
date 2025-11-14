package repo_team

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/sariya23/manage_pr_service/internal/outerror"
)

func (r *TeamRepository) GetTeam(ctx context.Context, teamName string) (string, error) {
	const operationPlace = "storage.repositories.team.GetTeam"

	getTeamSQL := fmt.Sprintf("select %s from %s where %s=$1", TeamTableTeamNameField, TeamTableName, TeamTableTeamNameField)
	var teamNameRes string
	err := r.conn.GetPool().QueryRow(ctx, getTeamSQL, teamName).Scan(&teamNameRes)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return "", fmt.Errorf("%s:%w", operationPlace, outerror.ErrTeamNotFound)
		}
		return "", fmt.Errorf("%s:%w", operationPlace, err)
	}

	return teamNameRes, nil
}
