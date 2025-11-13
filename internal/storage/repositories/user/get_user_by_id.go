package repo_user

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/sariya23/manage_pr_service/internal/models/domain"
	"github.com/sariya23/manage_pr_service/internal/outerror"
)

func (u *UserRepository) GetUserByID(ctx context.Context, userID int64) (*domain.User, error) {
	const operationPlace = "storage.repositories.user.GetUserByID"

	getUserSQL := fmt.Sprintf("select %s, %s, %s from '%s' where %s=$1",
		UserTableUserIDField,
		UserTableUsernameField,
		UserTableIsActiveField,
		UserTableName,
		UserTableUserIDField,
	)

	var user domain.User
	row := u.conn.GetPool().QueryRow(ctx, getUserSQL, userID)
	err := row.Scan(&user.UserID, &user.Username, &user.IsActive)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("%s: %w", operationPlace, outerror.ErrUserNotFound)
		} else {
			return nil, fmt.Errorf("%s: %w", operationPlace, err)
		}
	}
	return &user, nil
}
