package repo_team

import (
	"context"
	"fmt"
)

func (r *TeamRepository) GetTeamMemberIDs(ctx context.Context, teamName string) ([]string, error) {
	const operationPlace = "storage.repositories.team.GetTeamMemberIDs"

	getTeamMembersSQL := fmt.Sprintf("select %s from %s where %s=$1",
		TeamUserIDField,
		TeamTableName,
		TeamTeamNameField)

	var userIDs []string

	rows, err := r.conn.GetPool().Query(ctx, getTeamMembersSQL, teamName)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", operationPlace, err)
	}
	defer rows.Close()
	for rows.Next() {
		var userID string
		err = rows.Scan(&userID)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", operationPlace, err)
		}
		if rows.Err() != nil {
			return nil, fmt.Errorf("%s: %w", operationPlace, rows.Err())
		}
		userIDs = append(userIDs, userID)
	}
	return userIDs, nil
}
