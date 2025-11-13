package apiteams

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/go-chi/render"
	"github.com/sariya23/manage_pr_service/internal/converters"
	api "github.com/sariya23/manage_pr_service/internal/generated"
	"github.com/sariya23/manage_pr_service/internal/lib/erresponse"
	"github.com/sariya23/manage_pr_service/internal/outerror"
	teamsvalidators "github.com/sariya23/manage_pr_service/internal/validators/handlers/teams"
)

func (i *TeamsImplementation) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	const operationPlace = "handlers.users.GetReview"
	log := i.logger.With("operationPlace", operationPlace)
	query := r.URL.Query()
	teamName := query.Get("team_name")

	if msg, valid := teamsvalidators.ValidateTeamGet(teamName); !valid {
		log.Warn("invalid request", slog.String("team_name", teamName), slog.String("message", msg))
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, erresponse.MakeInvalidResponse(msg))
		return
	}

	members, err := i.teamsService.Get(ctx, teamName)
	if err != nil {
		if errors.Is(err, outerror.ErrTeamNotFound) {
			log.Warn("team not found", slog.String("team_name", teamName))
			w.WriteHeader(http.StatusNotFound)
			render.JSON(w, r, erresponse.MakeNotFoundResponse("resource not found"))
			return
		}
		log.Error("unexpected error", slog.String("team_name", teamName), slog.String("error", err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, erresponse.MakeInternalResponse("internal server error"))
		return
	}
	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, api.GetTeamGet200JSONResponse{
		TeamName: teamName,
		Members:  converters.MultiDomainUserToGetTeamResponse(members),
	})
	return
}
