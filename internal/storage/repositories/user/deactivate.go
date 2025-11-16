package user

import (
	"context"
	"fmt"
)

func (u *UserRepository) MultiDeactivate(ctx context.Context, userIDs []string) error {
	const operationPlace = "storage.repositories.user.MultiDeactivate"

	deactivateQuery := `update "user" set is_active=false where user_id=any($1)`

	_, err := u.conn.GetPool().Exec(ctx, deactivateQuery, userIDs)

	if err != nil {
		return fmt.Errorf("%s: %w", operationPlace, err)
	}

	return nil
}
