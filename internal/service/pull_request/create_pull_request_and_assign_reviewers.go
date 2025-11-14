package service_pull_request

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/sariya23/manage_pr_service/internal/lib/random"
	"github.com/sariya23/manage_pr_service/internal/models/domain"
	"github.com/sariya23/manage_pr_service/internal/models/dto"
	"github.com/sariya23/manage_pr_service/internal/outerror"
)

func (s *PullRequestService) CreatePullRequestAndAssignReviewers(ctx context.Context, prData dto.CreatePullRequestDTO) (*domain.PullRequest, []domain.User, error) {
	const operationPlace = "service.pull_request.create_pull_request_and_assign_reviewers"
	log := s.log.With("operation_place", operationPlace)

	_, err := s.PullRequestRepo.GetPullRequest(ctx, prData.ID)
	if err != nil {
		if !errors.Is(err, outerror.ErrPullRequestNotFound) {
			log.Error("unexpected error", slog.String("pr id", prData.ID), slog.String("error", err.Error()))
			return nil, nil, fmt.Errorf("%s:%w", operationPlace, err)
		}
	} else {
		log.Warn("pull request already exists", slog.String("pr id", prData.ID))
		return nil, nil, fmt.Errorf("%s:%w", operationPlace, outerror.ErrPullRequestAlreadyExists)
	}

	_, err = s.UserRepo.GetUserByID(ctx, prData.AuthorID)
	if err != nil {
		if errors.Is(err, outerror.ErrUserNotFound) {
			log.Warn("author does not exist", slog.String("user_id", prData.AuthorID))
			return nil, nil, fmt.Errorf("%s:%w", operationPlace, outerror.ErrUserNotFound)
		}
		log.Error("unexpected error while getting user", slog.String("user_id", prData.AuthorID), slog.String("error", err.Error()))
		return nil, nil, fmt.Errorf("%s:%w", operationPlace, err)
	}

	teamName, err := s.TeamRepo.GetUserTeam(ctx, prData.AuthorID)
	if err != nil {
		if errors.Is(err, outerror.ErrUserNotInAnyTeam) {
			log.Warn("user not in any team", slog.String("user_id", prData.AuthorID))
			return nil, nil, fmt.Errorf("%s: %w", operationPlace, outerror.ErrUserNotInAnyTeam)
		}
		log.Error("unexpected error while getting user team", slog.String("user_id", prData.AuthorID), slog.String("error", err.Error()))
		return nil, nil, fmt.Errorf("%s:%w", operationPlace, err)
	}

	teamMembers, err := s.TeamRepo.GetTeamMembers(ctx, teamName)
	if err != nil {
		log.Error("unexpected error while getting team members", slog.String("team_name", teamName), slog.String("error", err.Error()))
		return nil, nil, fmt.Errorf("%s:%w", operationPlace, err)
	}

	reviewers := pickReviewersForPR(onlyActiveUsers(teamMembers), prData.AuthorID)

	pullRequest, err := s.PullRequestRepo.CreatePullRequestAndAssignReviewers(ctx, prData, domain.UserIDs(reviewers))
	if err != nil {
		log.Error("cannot crate PR and assign reviewers", slog.String("pr_id", prData.ID), slog.String("error", err.Error()))
		return nil, nil, fmt.Errorf("%s: %w", operationPlace, err)
	}
	return pullRequest, teamMembers, nil
}

// onlyActiveUsers оставляет только активных юзеров
func onlyActiveUsers(users []domain.User) []domain.User {
	var result []domain.User

	for _, user := range users {
		if user.IsActive == true {
			result = append(result, user)
		}
	}
	return result
}

// pickReviewersForPR рандомно выбирает 2 или меньше ревьюера, исключая автора
func pickReviewersForPR(memberIDs []domain.User, authorID string) []domain.User {
	membersWithoutAuthor := make([]domain.User, 0, len(memberIDs)-1)
	for _, user := range memberIDs {
		if user.UserID != authorID {
			membersWithoutAuthor = append(membersWithoutAuthor, user)
		}
	}

	reviewers := random.Sample[domain.User](membersWithoutAuthor, 2)
	return reviewers
}
