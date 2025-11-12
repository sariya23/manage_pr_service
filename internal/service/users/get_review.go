package serviceusers

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/sariya23/manage_pr_service/internal/models"
	"github.com/sariya23/manage_pr_service/internal/outerror"
)

func (s *UsersService) GetReview(ctx context.Context, userID int64) ([]models.PullRequest, error) {
	const operationPlace = "service.users.GetReview"
	log := s.log.With("operationPlace", operationPlace)
	_, err := s.userRepo.GetUserByID(ctx, userID)
	if err != nil {
		if errors.Is(err, outerror.ErrUserNotFound) {
			return nil, fmt.Errorf("%s: %w", operationPlace, outerror.ErrUserNotFound)
		}
		log.Error("unexpected error while get user",
			slog.Int("user_id", int(userID)),
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("%s: %w", operationPlace, outerror.ErrInternal)
	}

	prs, err := s.reviewUserRepo.GetUserReviews(ctx, userID)
	if err != nil {
		log.Error("unexpected error while get user reviews", slog.Int64("user_id", userID), slog.String("error", err.Error()))
		return nil, fmt.Errorf("%s: %w", operationPlace, outerror.ErrInternal)
	}
	return prs, nil
}
