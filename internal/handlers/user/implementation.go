package apiuser

import "log/slog"

type userService interface {
}

type UsersImplementation struct {
	logger      *slog.Logger
	userService userService
}

func NewUsersImplementation(logger *slog.Logger, userService userService) *UsersImplementation {
	return &UsersImplementation{
		logger:      logger,
		userService: userService,
	}
}
