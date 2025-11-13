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
	apiusers "github.com/sariya23/manage_pr_service/internal/handlers/users"
	serviceusers "github.com/sariya23/manage_pr_service/internal/service/users"
	"github.com/sariya23/manage_pr_service/internal/storage/database"
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
	userRepo := repo_user.NewUserRepository(db, logger)
	teamRepo := repo_team.NewTeamRepository(db, logger)
	userService := serviceusers.NewUsersService(log, userRepo)
	debugImpl := apidebug.NewDebugImplementation()
	userImpl := apiusers.NewUsersImplementation(logger, userService)
	r.Route("/api", func(r chi.Router) {
		r.Route("/debug", func(r chi.Router) {
			r.Get("/ping", debugImpl.Ping)
		})
		r.Route("/users", func(r chi.Router) {
			r.Post("/setIsActive", userImpl.SetIsActive)
			r.Get("/getReview/{userID}", userImpl.GetReview)
		})
	})

	srv := server.NewServer(config.HTTPServerHost, config.HTTPServerPort, r)
	srv.MustRun()
}
