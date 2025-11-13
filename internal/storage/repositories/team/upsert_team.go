package repo_team

import (
	"context"
	"fmt"
	"strings"

	"github.com/sariya23/manage_pr_service/internal/models/domain"
	repo_user "github.com/sariya23/manage_pr_service/internal/storage/repositories/user"
)

func (r *TeamRepository) UpsertTeam(ctx context.Context, teamName string, users []domain.User) error {
	const operationPlace = "storage.repositories.team.UpsertTeam"

	upsertTeamSQL := fmt.Sprintf("insert into %s values ($1) on conflict (%s) do nothing",
		TeamTableName,
		TeamTableTeamNameField)

	insertTeamMemberSQL := fmt.Sprintf("insert into %s values ")
	insertTeamMemberArgs := make([]interface{}, 0, len(users)*2)
	insertTeamMemberValues := make([]string, 0, len(users))

	upsertUsersSQL := fmt.Sprintf("insert into \"%s\" values ")
	upsertUsersArgs := make([]interface{}, 0, len(users)*3)
	upsertUsersValues := make([]string, 0, len(users))

	for i, user := range users {
		upsertUsersArgs = append(upsertUsersArgs, user.UserID, user.Username, user.IsActive)
		upsertUsersValues = append(upsertUsersValues,
			fmt.Sprintf("($%d, $%d, $%d)",
				upsertUsersArgs[1+i],
				upsertUsersArgs[2+i],
				upsertUsersArgs[3+i]))

		insertTeamMemberArgs = append(insertTeamMemberArgs, teamName, user.UserID)
		insertTeamMemberValues = append(insertTeamMemberValues,
			fmt.Sprintf("($%d, $%d)",
				insertTeamMemberArgs[1+i],
				insertTeamMemberArgs[2+i]))
	}
	insertTeamMemberSQL += strings.Join(upsertUsersValues, ", ")
	upsertUsersSQL += strings.Join(upsertUsersValues, ", ")
	upsertUsersSQL += fmt.Sprintf("on conflict (%s) do nothing", repo_user.UserTableUserIDField)

	tx, err := r.conn.GetPool().Begin(ctx)
	if err != nil {
		return fmt.Errorf("%s: %w", operationPlace, err)
	}
	defer func() {
		err = tx.Rollback(ctx)
	}()

	_, err = tx.Exec(ctx, upsertTeamSQL, teamName)
	if err != nil {
		return fmt.Errorf("%s: %w", operationPlace, err)
	}

	_, err = tx.Exec(ctx, upsertUsersSQL, upsertUsersArgs)
	if err != nil {
		return fmt.Errorf("%s: %w", operationPlace, err)
	}

	_, err = tx.Exec(ctx, insertTeamMemberSQL, insertTeamMemberArgs)
	if err != nil {
		return fmt.Errorf("%s: %w", operationPlace, err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("%s: %w", operationPlace, err)
	}
	return nil
}
