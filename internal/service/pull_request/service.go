package service_pull_request

import "log/slog"

type PullRequestService struct {
	log *slog.Logger
}

func NewPullRequestService(log *slog.Logger) *PullRequestService {
	return &PullRequestService{
		log: log,
	}
}
