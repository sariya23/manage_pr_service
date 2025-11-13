package repo_user

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/sariya23/manage_pr_service/internal/models/domain"
	"github.com/sariya23/manage_pr_service/internal/outerror"
)

func (u *UserRepository) SetIsActive(ctx context.Context, userID int64, isActive bool) (*domain.User, error) {
	const operationPlace = "storage.repositories.user.SetIsActive"

	setIsActiveSQL := fmt.Sprintf("update '%s' set %s=$1 where %s=$2 returning %s, %s, %s",
		UserTableName,
		UserTableIsActiveField,
		UserTableUserIDField,
		UserTableUserIDField,
		UserTableUsernameField,
		UserTableIsActiveField,
	)
	var user domain.User
	row := u.conn.GetPool().QueryRow(ctx, setIsActiveSQL, userID, isActive)
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
