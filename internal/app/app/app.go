package app

import (
	"context"
	"log/slog"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/sariya23/manage_pr_service/internal/app/server"
	cfg "github.com/sariya23/manage_pr_service/internal/config"
	"github.com/sariya23/manage_pr_service/internal/handlers/analytics"
	apidebug "github.com/sariya23/manage_pr_service/internal/handlers/debug"
	api_pull_requests "github.com/sariya23/manage_pr_service/internal/handlers/pull_requests"
	apiteams "github.com/sariya23/manage_pr_service/internal/handlers/teams"
	apiusers "github.com/sariya23/manage_pr_service/internal/handlers/users"
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

	r := chi.NewRouter()
	r.Route("/api", func(r chi.Router) {
		r.Route("/debug", func(r chi.Router) {
			r.Get("/ping", debugImpl.Ping)
		})
		r.Route("/users", func(r chi.Router) {
			r.Post("/setIsActive", userImpl.SetIsActive)
			r.Get("/getReview/{user_id}", userImpl.GetReview)
		})
		r.Route("/team", func(r chi.Router) {
			r.Post("/add", teamsImpl.Add)
			r.Get("/get/{team_name}", teamsImpl.Get)
		})
		r.Route("/pullRequest", func(r chi.Router) {
			r.Post("/create", pullRequestImpl.Create)
			r.Post("/merge", pullRequestImpl.Merge)
			r.Post("/reassign", pullRequestImpl.Reassign)
		})
		r.Route("/analytics", func(r chi.Router) {
			r.Get("/usersPRs", analyticsImpl.UsersPRs)
		})
	})
	srv := server.NewServer(config.HTTPServerHost, config.HTTPServerPort, r)
	return &App{srv: srv}
}

func (a *App) MustStart() {
	a.srv.MustRun()
}
func (a *App) GracefulStop(ctx context.Context) {
	a.srv.Srv.Shutdown(ctx)

}
