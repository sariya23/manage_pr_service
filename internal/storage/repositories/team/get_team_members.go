package repo_team

import (
	"context"
	"fmt"

	"github.com/sariya23/manage_pr_service/internal/models/domain"
	repo_user "github.com/sariya23/manage_pr_service/internal/storage/repositories/user"
)

func (r *TeamRepository) GetTeamMembers(ctx context.Context, teamName string) ([]domain.User, error) {
	const operationPlace = "storage.repositories.team.GetTeamMember"

	getTeamMembersSQL := fmt.Sprintf("select %s, %s, %s from %s join %s using(%s) join \"%s\" using(%s) where %s=$1",
		TeamMemberTableUserIDField,
		repo_user.UserTableUsernameField,
		repo_user.UserTableIsActiveField,
		TeamTableName,
		TeamMemberTableName,
		TeamTableTeamNameField,
		repo_user.UserTableName,
		repo_user.UserTableUserIDField,
		TeamTableTeamNameField)
	var users []domain.User

	rows, err := r.conn.GetPool().Query(ctx, getTeamMembersSQL, teamName)
	if err != nil {
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
