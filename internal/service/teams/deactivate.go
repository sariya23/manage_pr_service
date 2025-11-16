package serviceteams

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"slices"
	"strings"

	"github.com/sariya23/manage_pr_service/internal/lib/request"
	"github.com/sariya23/manage_pr_service/internal/models/domain"
	"github.com/sariya23/manage_pr_service/internal/outerror"
)

func (s *TeamsService) Deactivate(ctx context.Context, teamName string, userIDs []string) error {
	const operationPlace = "service.teams.Deactivate"
	log := s.log.With("operationPlace", operationPlace)
	requestID := request.GetIDKey(ctx)
	log = log.With("request_id", requestID)

	_, err := s.teamRepository.GetTeam(ctx, teamName)
	if err != nil {
		if errors.Is(err, outerror.ErrTeamNotFound) {
			log.Warn("team not found", slog.String("team_name", teamName))
			return fmt.Errorf("%s:%w", operationPlace, err)
		}
		log.Error("failed to get team", slog.String("team_name", teamName), slog.String("error", err.Error()))
		return fmt.Errorf("%s:%w", operationPlace, err)
	}

	teamMembers, err := s.teamRepository.GetTeamMembers(ctx, teamName)
	if err != nil {
		log.Error("failed to get team members", slog.String("team_name", teamName))
		return fmt.Errorf("%s:%w", operationPlace, err)
	}

	teamMemberIDs := domain.UserIDs(teamMembers)
	for _, userID := range userIDs {
		if !slices.Contains(teamMemberIDs, userID) {
			log.Warn("team member not found", slog.String("team_name", teamName), slog.String("user_id", userID))
			return fmt.Errorf("%s:%w", operationPlace, outerror.ErrUserNotInTeam)
		}
	}

	err = s.userRepository.MultiDeactivate(ctx, userIDs)
	if err != nil {
		log.Error("failed to multi deactivate users", slog.String("user_ids", strings.Join(userIDs, ",")), slog.String("error", err.Error()))
		return fmt.Errorf("%s: %w", operationPlace, err)
	}
	return nil
}
