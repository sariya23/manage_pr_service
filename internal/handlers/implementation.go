package handlers

import (
	"github.com/sariya23/manage_pr_service/internal/handlers/analytics"
	apidebug "github.com/sariya23/manage_pr_service/internal/handlers/debug"
	api_pull_requests "github.com/sariya23/manage_pr_service/internal/handlers/pull_requests"
	apiteams "github.com/sariya23/manage_pr_service/internal/handlers/teams"
	apiusers "github.com/sariya23/manage_pr_service/internal/handlers/users"
)

type Implementation struct {
	apidebug.DebugImplementation
	analytics.AnalyticsImplementation
	apiusers.UsersImplementation
	apiteams.TeamsImplementation
	api_pull_requests.PullRequestImplementation
}

func NewImplementation(debugImpl apidebug.DebugImplementation,
	analyticsImpl analytics.AnalyticsImplementation,
	usersImpl apiusers.UsersImplementation,
	teamsImpl apiteams.TeamsImplementation,
	pullRequestImpl api_pull_requests.PullRequestImplementation) Implementation {
	return Implementation{
		debugImpl,
		analyticsImpl, usersImpl,
		teamsImpl,
		pullRequestImpl,
	}

}
