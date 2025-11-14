package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/sariya23/manage_pr_service/internal/app/server"
	cfg "github.com/sariya23/manage_pr_service/internal/config"
	apidebug "github.com/sariya23/manage_pr_service/internal/handlers/debug"
	api_pull_requests "github.com/sariya23/manage_pr_service/internal/handlers/pull_requests"
	apiteams "github.com/sariya23/manage_pr_service/internal/handlers/teams"
	apiusers "github.com/sariya23/manage_pr_service/internal/handlers/users"
	service_pull_request "github.com/sariya23/manage_pr_service/internal/service/pull_request"
	serviceteams "github.com/sariya23/manage_pr_service/internal/service/teams"
	serviceusers "github.com/sariya23/manage_pr_service/internal/service/users"
	"github.com/sariya23/manage_pr_service/internal/storage/database"
	repo_pull_request "github.com/sariya23/manage_pr_service/internal/storage/repositories/pull_request"
	repo_team "github.com/sariya23/manage_pr_service/internal/storage/repositories/team"
	repo_user "github.com/sariya23/manage_pr_service/internal/storage/repositories/user"
)

func main() {
	ctx := context.Background()
	config := cfg.MustLoad()
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	fmt.Println(config)

	r := chi.NewRouter()
	dbURL := database.GenerateDBUrl(
		config.PostgresUsername,
		config.PostgresPassword,
		config.PostgresInnerHost,
		strconv.Itoa(config.PostgresPort),
		config.PostgresDB,
		config.SSLMode)
	db := database.MustNewConnection(ctx, logger, dbURL)
	userRepo := repo_user.NewUserRepository(db)
	teamRepo := repo_team.NewTeamRepository(db)
	pullRequestRepo := repo_pull_request.NewPullRequestRepository(db)
	userService := serviceusers.NewUsersService(logger, userRepo, pullRequestRepo, teamRepo)
	teamService := serviceteams.NewTeamsService(logger, teamRepo, userRepo)
	pullRequestService := service_pull_request.NewPullRequestService(logger, pullRequestRepo, userRepo, teamRepo)
	debugImpl := apidebug.NewDebugImplementation()
	userImpl := apiusers.NewUsersImplementation(logger, userService)
	teamsImpl := apiteams.NewTeamsImplementation(logger, teamService)
	pullRequestImpl := api_pull_requests.NewPullRequestImplementation(logger, pullRequestService)
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
		})
	})

	srv := server.NewServer(config.HTTPServerHost, config.HTTPServerPort, r)
	srv.MustRun()
}
