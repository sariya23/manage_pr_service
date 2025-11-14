package serviceusers

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/sariya23/manage_pr_service/internal/outerror"
)

func (s *UsersService) GetUserTeam(ctx context.Context, userID string) (string, error) {
	const operationPlace = "service.users.GetUserTeam"
	log := s.log.With("operationPlace", operationPlace)

	teamName, err := s.teamRepo.GetUserTeam(ctx, userID)
	if err != nil {
		if errors.Is(err, outerror.ErrUserNotInAnyTeam) {
			log.Warn("user not found", slog.String("user_id", userID))
			return "", fmt.Errorf("%s: %w", operationPlace, outerror.ErrUserNotFound)
		}
		log.Error("failed to get user team",
			slog.String("user_id", userID),
			slog.String("error", err.Error()))
		return "", fmt.Errorf("%s: %w", operationPlace, err)
	}
	return teamName, nil
}
