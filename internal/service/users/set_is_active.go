package serviceusers

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/sariya23/manage_pr_service/internal/models/domain"
	"github.com/sariya23/manage_pr_service/internal/outerror"
)

func (s *UsersService) SetIsActive(ctx context.Context, userId string, isActive bool) (*domain.User, error) {
	const operationPlace = "service.users.SetIsActive"
	log := s.log.With("operationPlace", operationPlace)
	user, err := s.userRepo.SetIsActive(ctx, userId, isActive)
	if err != nil {
		if errors.Is(err, outerror.ErrUserNotFound) {
			return nil, fmt.Errorf("%s: %w", operationPlace, outerror.ErrUserNotFound)
		}
		log.Error("unexpected error while update is_active",
			slog.String("user_id", userId),
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("%s: %w", operationPlace, outerror.ErrInternal)
	}
	return user, nil
}
