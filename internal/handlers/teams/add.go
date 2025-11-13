package apiteams

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/go-chi/render"
	"github.com/sariya23/manage_pr_service/internal/converters"
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

	members, err := i.teamsService.Add(ctx, request.TeamName, converters.MultiAddTeamUserToDomainUser(request.Members))
	if err != nil {
		if errors.Is(err, outerror.ErrTeamAlreadyExists) {
			log.Warn("team already exists", slog.String("teamname", request.TeamName))
			w.WriteHeader(http.StatusBadRequest)
			render.JSON(w, r, erresponse.MakeTeamAlreadyExistsResponse(fmt.Sprintf("%s already exists", request.TeamName)))
			return
		} else if errors.Is(err, outerror.ErrUserAlreadyInTeam) {
			log.Warn("user already in team", slog.String("teamname", request.TeamName))
			w.WriteHeader(http.StatusBadRequest)
			render.JSON(w, r, erresponse.MakeInvalidResponse(fmt.Sprintf("one of users already in team %s", request.TeamName)))
			return
		} else if errors.Is(err, outerror.ErrInactiveUser) {
			log.Warn("one of user is inactive", slog.String("teamname", request.TeamName))
			w.WriteHeader(http.StatusBadRequest)
			render.JSON(w, r, erresponse.MakeInvalidResponse("one of users is inactive"))
			return
		}
		log.Error("unexpected error", slog.String("error", err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, erresponse.MakeInternalResponse("internal server error"))
		return
	}
	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, api.PostTeamAdd201JSONResponse{Team: &api.Team{
		Members:  converters.MultiDomainUserToAddTeamResponse(members),
		TeamName: request.TeamName,
	}})
}
