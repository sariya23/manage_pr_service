package repo_team

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/sariya23/manage_pr_service/internal/outerror"
)

func (r *TeamRepository) GetUserTeam(ctx context.Context, userID string) (string, error) {
	const operationPlace = "storage.repositories.team.GetUserTeam"

	getUserTeamSQL := fmt.Sprintf("select %s from %s join %s using(%s) where %s=$1",
		TeamTableTeamNameField,
		TeamTableName,
		TeamMemberTableName,
		TeamTableTeamNameField,
		TeamMemberTableUserIDField,
	)

	var teamName string
	row := r.conn.GetPool().QueryRow(ctx, getUserTeamSQL, userID)
	err := row.Scan(teamName)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return "", fmt.Errorf("%s: %w", operationPlace, outerror.ErrUserNotFound)
		} else {
			return "", fmt.Errorf("%s: %w", operationPlace, err)
		}
	}

	return teamName, nil
}
