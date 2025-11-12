package main

import (
	"fmt"

	"github.com/go-chi/chi/v5"
	"github.com/sariya23/manage_pr_service/internal/app/server"
	cfg "github.com/sariya23/manage_pr_service/internal/config"
	apidebug "github.com/sariya23/manage_pr_service/internal/handlers/debug"
)

func main() {
	config := cfg.MustLoad()
	fmt.Println(config)

	r := chi.NewRouter()
	debugImpl := apidebug.NewDebugImplementation()
	r.Get("/api/ping", debugImpl.Ping)

	srv := server.NewServer(config.HTTPServerHost, config.HTTPServerPort, r)
	srv.MustRun()
}
