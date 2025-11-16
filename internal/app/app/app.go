package app

import (
	"context"
	"log/slog"
	"strconv"

	"github.com/sariya23/manage_pr_service/internal/app/server"
	cfg "github.com/sariya23/manage_pr_service/internal/config"
	api "github.com/sariya23/manage_pr_service/internal/generated"
	"github.com/sariya23/manage_pr_service/internal/handlers"
	"github.com/sariya23/manage_pr_service/internal/handlers/analytics"
	apidebug "github.com/sariya23/manage_pr_service/internal/handlers/debug"
	api_pull_requests "github.com/sariya23/manage_pr_service/internal/handlers/pull_requests"
	apiteams "github.com/sariya23/manage_pr_service/internal/handlers/teams"
	apiusers "github.com/sariya23/manage_pr_service/internal/handlers/users"
	"github.com/sariya23/manage_pr_service/internal/middleware"
	serviceanalytics "github.com/sariya23/manage_pr_service/internal/service/analytics"
	service_pull_request "github.com/sariya23/manage_pr_service/internal/service/pull_request"
	serviceteams "github.com/sariya23/manage_pr_service/internal/service/teams"
	serviceusers "github.com/sariya23/manage_pr_service/internal/service/users"
	"github.com/sariya23/manage_pr_service/internal/storage/database"
	repo_pull_request "github.com/sariya23/manage_pr_service/internal/storage/repositories/pull_request"
	repo_team "github.com/sariya23/manage_pr_service/internal/storage/repositories/team"
	repo_user "github.com/sariya23/manage_pr_service/internal/storage/repositories/user"
)

type App struct {
	srv *server.Server
}

func NewApp(ctx context.Context, logger *slog.Logger, config *cfg.Config) *App {
	dbURL := database.GenerateDBUrl(
		config.PostgresUsername,
		config.PostgresPassword,
		config.PostgresInnerHost,
		strconv.Itoa(config.PostgresPort),
		config.PostgresDB,
		config.SSLMode)
	db := database.MustNewConnection(ctx, logger, dbURL)

	// Repos
	userRepo := repo_user.NewUserRepository(db)
	teamRepo := repo_team.NewTeamRepository(db)
	pullRequestRepo := repo_pull_request.NewPullRequestRepository(db)

	// Services
	userService := serviceusers.NewUsersService(logger, userRepo, pullRequestRepo, teamRepo)
	teamService := serviceteams.NewTeamsService(logger, teamRepo, userRepo)
	pullRequestService := service_pull_request.NewPullRequestService(logger, pullRequestRepo, userRepo, teamRepo)
	analyticsService := serviceanalytics.NewAnalyticsService(logger, pullRequestRepo)

	// Implementations
	debugImpl := apidebug.NewDebugImplementation()
	userImpl := apiusers.NewUsersImplementation(logger, userService)
	teamsImpl := apiteams.NewTeamsImplementation(logger, teamService)
	pullRequestImpl := api_pull_requests.NewPullRequestImplementation(logger, pullRequestService)
	analyticsImpl := analytics.NewAnalyticsImplementation(logger, analyticsService)
	impl := handlers.NewImplementation(debugImpl, analyticsImpl, userImpl, teamsImpl, pullRequestImpl)

	router := api.HandlerWithOptions(impl, api.ChiServerOptions{BaseURL: "/api", Middlewares: []api.MiddlewareFunc{middleware.RequestIDMiddleware}})
	srv := server.NewServer(config.HTTPServerHost, config.HTTPServerPort, router)
	return &App{srv: srv}
}

func (a *App) MustStart() {
	a.srv.MustRun()
}
func (a *App) GracefulStop(ctx context.Context) {
	a.srv.Srv.Shutdown(ctx)

}
