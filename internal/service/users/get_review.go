package serviceusers

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/sariya23/manage_pr_service/internal/models/domain"
	"github.com/sariya23/manage_pr_service/internal/outerror"
)

func (s *UsersService) GetReviews(ctx context.Context, userID string) ([]domain.PullRequest, error) {
	const operationPlace = "service.users.GetReview"
	log := s.log.With("operationPlace", operationPlace)
	_, err := s.userRepo.GetUserByID(ctx, userID)
	if err != nil {
		if errors.Is(err, outerror.ErrUserNotFound) {
			return nil, fmt.Errorf("%s: %w", operationPlace, outerror.ErrUserNotFound)
		}
		log.Error("unexpected error while get user",
			slog.String("user_id", userID),
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("%s: %w", operationPlace, outerror.ErrInternal)
	}

	prs, err := s.reviewUserRepo.GetUserReviews(ctx, userID)
	if err != nil {
		log.Error("unexpected error while get user reviews", slog.String("user_id", userID), slog.String("error", err.Error()))
		return nil, fmt.Errorf("%s: %w", operationPlace, outerror.ErrInternal)
	}
	return prs, nil
}
