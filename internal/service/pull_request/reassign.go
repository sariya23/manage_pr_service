package service_pull_request

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"slices"

	"github.com/sariya23/manage_pr_service/internal/lib/random"
	"github.com/sariya23/manage_pr_service/internal/models/domain"
	"github.com/sariya23/manage_pr_service/internal/outerror"
)

func (s *PullRequestService) Reassign(ctx context.Context, prID string, oldReviewerID string) (*domain.PullRequest, string, error) {
	const operationPlace = "service.pull_request.Reassign"
	log := s.log.With("operationPlace", operationPlace)

	pr, err := s.PullRequestRepo.GetPullRequest(ctx, prID)
	if err != nil {
		if errors.Is(err, outerror.ErrPullRequestNotFound) {
			return nil, "", fmt.Errorf("%s:%w", operationPlace, err)
		}
		log.Error("failed to get pull request", slog.String("pr id", prID), slog.String("error", err.Error()))
		return nil, "", fmt.Errorf("%s:%w", operationPlace, err)
	}
	_, err = s.UserRepo.GetUserByID(ctx, oldReviewerID)
	if err != nil {
		if errors.Is(err, outerror.ErrUserNotFound) {
			return nil, "", fmt.Errorf("%s:%w", operationPlace, err)
		}
		log.Error("failed to get user", slog.String("user_id", oldReviewerID), slog.String("error", err.Error()))
		return nil, "", fmt.Errorf("%s:%w", operationPlace, err)
	}

	if pr.AuthorID == oldReviewerID {
		log.Warn("author is not reviewer", slog.String("pr_id", prID))
		return nil, "", fmt.Errorf("%s:%w", operationPlace, outerror.ErrUserIsNotReviewer)
	}

	teamName, err := s.TeamRepo.GetUserTeam(ctx, pr.AuthorID)
	if err != nil {
		if errors.Is(err, outerror.ErrUserNotInAnyTeam) {
			log.Error("PR author is not in any team", slog.String("pr_id", prID), slog.String("author_id", pr.AuthorID))
			return nil, "", fmt.Errorf("%s:%w", operationPlace, err)
		}
		log.Error("failed to get team", slog.String("pr_id", prID), slog.String("error", err.Error()))
		return nil, "", fmt.Errorf("%s:%w", operationPlace, err)
	}
	members, err := s.TeamRepo.GetTeamMembers(ctx, teamName)
	if err != nil {
		log.Error("cannot get team members", slog.String("team name", teamName), slog.String("error", err.Error()))
		return nil, "", fmt.Errorf("%s:%w", operationPlace, err)
	}
	if !slices.Contains(domain.UserIDs(members), oldReviewerID) {
		log.Warn("user not in PR team", slog.String("pr_id", prID),
			slog.String("old_reviewer_id", oldReviewerID),
			slog.String("team name", teamName))
		return nil, "", fmt.Errorf("%s:%w", operationPlace, outerror.ErrUserNotInPullRequestTeam)
	}

	members = excludePRAuthorFromTeamMembers(onlyActiveUsers(members), pr.AuthorID)

	if len(members) == 0 {
		log.Warn("no reviewer candidates", slog.String("pr_id", prID),
			slog.String("author_id", pr.AuthorID), slog.String("tema name", teamName))
		return nil, "", fmt.Errorf("%s:%w", operationPlace, outerror.ErrNoReviewerCandidates)
	}
	newReviewerID := random.Choice(members)

	updatedPR, err := s.PullRequestRepo.ReassignPullRequest(ctx, prID, oldReviewerID, newReviewerID.UserID)
	if err != nil {
		log.Error("failed to reassign PR", slog.String("pr_id", prID), slog.String("error", err.Error()))
		return nil, "", fmt.Errorf("%s:%w", operationPlace, err)
	}
	return updatedPR, newReviewerID.UserID, nil
}

func excludePRAuthorFromTeamMembers(members []domain.User, authorID string) []domain.User {
	res := make([]domain.User, 0, len(members)-1)
	for _, m := range members {
		if m.UserID != authorID {
			res = append(res, m)
		}
	}
	return res
}
