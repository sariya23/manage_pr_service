package team

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

	insertTeamMemberSQL := strings.Builder{}
	insertTeamMemberSQL.WriteString(fmt.Sprintf("insert into %s values ", TeamMemberTableName))
	insertTeamMemberArgs := make([]interface{}, 0, len(users)*2)
	insertTeamMemberValues := make([]string, 0, len(users))

	insertUsersSQL := strings.Builder{}
	insertUsersSQL.WriteString(fmt.Sprintf("insert into \"%s\" values ", repo_user.UserTableName))
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
	insertTeamMemberSQL.WriteString(strings.Join(insertTeamMemberValues, ", "))
	insertUsersSQL.WriteString(strings.Join(insertUsersValues, ", "))
	insertUsersSQL.WriteString(fmt.Sprintf(" on conflict (%s) do nothing", repo_user.UserTableUserIDField))

	tx, err := r.conn.GetPool().Begin(ctx)
	if err != nil {
		return fmt.Errorf("%s: %w", operationPlace, err)
	}
	defer func() {
		err = tx.Rollback(ctx)
	}()

	_, err = tx.Exec(ctx, insertTeamSQL, teamName)
	if err != nil {
		return fmt.Errorf("%s:insertTeamSQL: %w", operationPlace, err)
	}

	if len(users) > 0 {
		_, err = tx.Exec(ctx, insertUsersSQL.String(), insertUsersArgs...)
		if err != nil {
			return fmt.Errorf("%s:insertUsersSQL: %w", operationPlace, err)
		}
	}

	if len(users) > 0 {
		_, err = tx.Exec(ctx, insertTeamMemberSQL.String(), insertTeamMemberArgs...)
		if err != nil {
			return fmt.Errorf("%s:insertTeamMemberSQL: %w", operationPlace, err)
		}
	}

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("%s: %w", operationPlace, err)
	}
	return nil
}
