package team

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/sariya23/manage_pr_service/internal/models/domain"
)

func (r *TeamRepository) GetTeamMembers(ctx context.Context, teamName string) ([]domain.User, error) {
	const operationPlace = "storage.repositories.team.GetTeamMember"

	getTeamMembersSQL := `select user_id, username, is_active from team join team_member using(team_name) join "user"
using(user_id) where team_name=$1`
	var users []domain.User

	rows, err := r.conn.GetPool().Query(ctx, getTeamMembersSQL, teamName)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return users, nil
		}
		return nil, fmt.Errorf("%s: %w", operationPlace, err)
	}
	defer rows.Close()
	for rows.Next() {
		var user domain.User
		err = rows.Scan(&user.UserID, &user.Username, &user.IsActive)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", operationPlace, err)
		}
		if rows.Err() != nil {
			return nil, fmt.Errorf("%s: %w", operationPlace, rows.Err())
		}
		users = append(users, user)
	}
	return users, nil
}
