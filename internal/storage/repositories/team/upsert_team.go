package repo_team

import (
	"context"
	"fmt"
	"strings"

	"github.com/sariya23/manage_pr_service/internal/models/domain"
	repo_user "github.com/sariya23/manage_pr_service/internal/storage/repositories/user"
)

func (r *TeamRepository) InsertTeam(ctx context.Context, teamName string, users []domain.User) error {
	const operationPlace = "storage.repositories.team.UpsertTeam"

	insertTeamSQL := fmt.Sprintf("insert into %s values ($1) on conflict (%s) do nothing",
		TeamTableName,
		TeamTableTeamNameField)

	insertTeamMemberSQL := fmt.Sprintf("insert into %s values ", TeamMemberTableName)
	insertTeamMemberArgs := make([]interface{}, 0, len(users)*2)
	insertTeamMemberValues := make([]string, 0, len(users))

	insertUsersSQL := fmt.Sprintf("insert into \"%s\" values ", repo_user.UserTableName)
	insertUsersArgs := make([]interface{}, 0, len(users)*3)
	insertUsersValues := make([]string, 0, len(users))

	for _, user := range users {
		insertUsersValues = append(insertUsersValues,
			fmt.Sprintf("($%d, $%d, $%d)",
				len(insertUsersArgs)+1,
				len(insertUsersArgs)+2,
				len(insertUsersArgs)+3))
		insertUsersArgs = append(insertUsersArgs, user.UserID, user.Username, user.IsActive)
		insertTeamMemberValues = append(insertTeamMemberValues,
			fmt.Sprintf("($%d, $%d)",
				len(insertTeamMemberArgs)+1,
				len(insertTeamMemberArgs)+2))
		insertTeamMemberArgs = append(insertTeamMemberArgs, teamName, user.UserID)
	}
	insertTeamMemberSQL += strings.Join(insertTeamMemberValues, ", ")
	insertUsersSQL += strings.Join(insertUsersValues, ", ")
	insertUsersSQL += fmt.Sprintf(" on conflict (%s) do nothing", repo_user.UserTableUserIDField)

	tx, err := r.conn.GetPool().Begin(ctx)
	if err != nil {
		return fmt.Errorf("%s: %w", operationPlace, err)
	}
	defer func() {
		err = tx.Rollback(ctx)
	}()

	_, err = tx.Exec(ctx, insertTeamSQL, teamName)
	if err != nil {
		fmt.Println("upsertTeamSQL")
		return fmt.Errorf("%s: %w", operationPlace, err)
	}

	_, err = tx.Exec(ctx, insertUsersSQL, insertUsersArgs...)
	if err != nil {
		fmt.Println("upsertUsersSQL")
		return fmt.Errorf("%s: %w", operationPlace, err)
	}

	_, err = tx.Exec(ctx, insertTeamMemberSQL, insertTeamMemberArgs...)
	if err != nil {
		fmt.Println("insertTeamMemberSQL")
		return fmt.Errorf("%s: %w", operationPlace, err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("%s: %w", operationPlace, err)
	}
	return nil
}
