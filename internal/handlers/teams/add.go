package apiteams

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/go-chi/render"
	teamsconverters "github.com/sariya23/manage_pr_service/internal/converters/handlers/teams"
	api "github.com/sariya23/manage_pr_service/internal/generated"
	"github.com/sariya23/manage_pr_service/internal/outerror"
	"github.com/sariya23/manage_pr_service/internal/utils/erresponse"
	teamsvalidators "github.com/sariya23/manage_pr_service/internal/validators/handlers/teams"
)

func (i *TeamsImplementation) Add(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	const operationPlace = "handlers.teams.Add"
	log := i.logger.With("operationPlace", operationPlace)

	var request api.PostTeamAddJSONRequestBody
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		log.Error("error decoding request body", slog.String("error", err.Error()))
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, erresponse.MakeInvalidResponse("invalid json"))
		return
	}
	defer func() {
		if err := r.Body.Close(); err != nil {
			log.Warn("error closing body", slog.String("error", err.Error()))
		}
	}()

	if msg, valid := teamsvalidators.ValidateTeamAddRequest(request); !valid {
		log.Warn("invalid request", slog.String("message", msg))
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, erresponse.MakeInvalidResponse(msg))
		return
	}

	_, err := i.teamsService.Add(ctx, request.TeamName, teamsconverters.MultiToDTOMember(request.Members))
	if err != nil {
		if errors.Is(err, outerror.ErrTeamAlreadyExists) {
			log.Warn("team already exists", slog.String("teamname", request.TeamName))
			w.WriteHeader(http.StatusBadRequest)
			render.JSON(w, r, erresponse.MakeTeamAlreadyExistsResponse(fmt.Sprintf("%s already exists", request.TeamName)))
			return
		} else if errors.Is(err, outerror.ErrTeamAlreadyExists) {
			log.Warn("user already in team", slog.String("teamname", request.TeamName))
			w.WriteHeader(http.StatusBadRequest)
			render.JSON(w, r, erresponse.MakeInvalidResponse(fmt.Sprintf("one of users already in team %s", request.TeamName)))
			return
		}
	}
}
