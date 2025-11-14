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

	getUserTeamSQL := `select team_name from team join team_member using(team_name where user_id=$1)`
	var teamName string
	fmt.Println(getUserTeamSQL)
	row := r.conn.GetPool().QueryRow(ctx, getUserTeamSQL, userID)
	err := row.Scan(&teamName)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return "", fmt.Errorf("%s: %w", operationPlace, outerror.ErrUserNotInAnyTeam)
		}
		return "", fmt.Errorf("%s: %w", operationPlace, err)
	}

	return teamName, nil
}
