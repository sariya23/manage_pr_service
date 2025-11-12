package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/sariya23/manage_pr_service/internal/app/server"
	cfg "github.com/sariya23/manage_pr_service/internal/config"
	apidebug "github.com/sariya23/manage_pr_service/internal/handlers/debug"
	apiusers "github.com/sariya23/manage_pr_service/internal/handlers/users"
	serviceusers "github.com/sariya23/manage_pr_service/internal/service/users"
)

func main() {
	config := cfg.MustLoad()
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	fmt.Println(config)

	r := chi.NewRouter()
	userService := serviceusers.NewUsersService()
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
