package outerror

import "errors"

var (
	ErrUserNotFound      = errors.New("user not found")
	ErrInternal          = errors.New("internal error")
	ErrTeamAlreadyExists = errors.New("team already exists")
	ErrUserAlreadyInTeam = errors.New("user already in team")
)
