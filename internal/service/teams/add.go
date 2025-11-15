package serviceteams

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"slices"

	"github.com/sariya23/manage_pr_service/internal/models/domain"
	"github.com/sariya23/manage_pr_service/internal/outerror"
)

func (s *TeamsService) Add(ctx context.Context, teamName string, membersRequest []domain.User) ([]domain.User, error) {
	const operationPlace = "service.teams.Add"
	log := s.log.With("operationPlace", operationPlace)
	members, err := s.teamRepository.GetTeamMembers(ctx, teamName)
	if err != nil {
		log.Error("failed to check team existence",
			slog.String("teamname", teamName),
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("%s: %w", operationPlace, err)

	}

	if len(members) > 0 {
		for _, memberReq := range membersRequest {
			if slices.Contains(domain.UserIDs(members), memberReq.UserID) {
				log.Warn("team already exists", slog.String("teamname", teamName))
				return nil, fmt.Errorf("%s: %w", operationPlace, outerror.ErrTeamAlreadyExists)
			}
		}
	}

	for _, member := range membersRequest {
		if !member.IsActive {
			log.Warn("member is not active", slog.String("user_id", member.UserID))
			return nil, fmt.Errorf("%s: %w", operationPlace, outerror.ErrInactiveUser)
		}
		_, err := s.userRepository.GetUserByID(ctx, member.UserID)
		isUserExists := true
		if err != nil {
			if errors.Is(err, outerror.ErrUserNotFound) {
				isUserExists = false
			} else {
				log.Error("failed to check user existence",
					slog.String("user_id", member.UserID),
					slog.String("error", err.Error()))
				return nil, fmt.Errorf("%s: %w", operationPlace, err)
			}

		}
		if isUserExists {
			_, err := s.teamRepository.GetUserTeam(ctx, member.UserID)
			isUserInTeam := true
			if err != nil {
				if errors.Is(err, outerror.ErrUserNotInAnyTeam) {
					isUserInTeam = false
				} else {
					log.Error("failed to check user membership in team",
						slog.String("user_id", member.UserID),
						slog.String("error", err.Error()))
					return nil, fmt.Errorf("%s: %w", operationPlace, err)
				}
			}
			if isUserInTeam {
				log.Warn("user is already a member of some team", slog.String("user_id", member.UserID))
				return nil, fmt.Errorf("%s: %w", operationPlace, outerror.ErrUserAlreadyInTeam)
			}
		}
	}

	err = s.teamRepository.InsertTeam(ctx, teamName, membersRequest)
	if err != nil {
		log.Error("failed to upsert team", slog.String("teamname", teamName), slog.String("error", err.Error()))
		return nil, fmt.Errorf("%s: %w", operationPlace, err)
	}
	membersUpdated, err := s.teamRepository.GetTeamMembers(ctx, teamName)
	if err != nil {
		log.Error("failed to get team members", slog.String("teamname", teamName), slog.String("error", err.Error()))
		return nil, fmt.Errorf("%s: %w", operationPlace, err)
	}
	return membersUpdated, nil
}
