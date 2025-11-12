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
	serviceuser "github.com/sariya23/manage_pr_service/internal/service/user"
)

func main() {
	config := cfg.MustLoad()
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	fmt.Println(config)

	r := chi.NewRouter()
	userService := serviceuser.NewUserService()
	debugImpl := apidebug.NewDebugImplementation()
	userImpl := apiusers.NewUsersImplementation(logger, userService)
	r.Route("/api", func(r chi.Router) {
		r.Route("/debug", func(r chi.Router) {
			r.Get("/ping", debugImpl.Ping)
		})
		r.Route("/user", func(r chi.Router) {
			r.Post("/setIsActive", userImpl.SetIsActive)
		})
	})

	srv := server.NewServer(config.HTTPServerHost, config.HTTPServerPort, r)
	srv.MustRun()
}
