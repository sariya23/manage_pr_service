package outerror

import "errors"

var (
	ErrUserNotFound             = errors.New("user not found")
	ErrInternal                 = errors.New("internal error")
	ErrTeamAlreadyExists        = errors.New("team already exists")
	ErrUserAlreadyInTeam        = errors.New("user already in team")
	ErrInactiveUser             = errors.New("inactive user")
	ErrTeamNotFound             = errors.New("team not found")
	ErrUserNotInAnyTeam         = errors.New("user not in any team")
	ErrPullRequestAlreadyExists = errors.New("pull request already exists")
	ErrPullRequestNotFound      = errors.New("pull request not found")
	ErrPullRequestMerged        = errors.New("pull request merged")
	ErrUserIsNotReviewer        = errors.New("user is not reviewer")
	ErrNoReviewerCandidates     = errors.New("no reviewers candidates")
)
