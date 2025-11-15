package apiteams

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/go-chi/render"
	"github.com/sariya23/manage_pr_service/internal/converters"
	api "github.com/sariya23/manage_pr_service/internal/generated"
	"github.com/sariya23/manage_pr_service/internal/lib/erresponse"
	"github.com/sariya23/manage_pr_service/internal/lib/errorhandler"
	teamsvalidators "github.com/sariya23/manage_pr_service/internal/validators"
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
	if status, resp, isError := errorhandler.TeamAdd(err, request.TeamName); isError {
		w.WriteHeader(status)
		render.JSON(w, r, resp)
		return
	}
	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, api.PostTeamAdd201JSONResponse{Team: &api.Team{
		Members:  converters.MultiDomainUserToAddTeamResponse(members),
		TeamName: request.TeamName,
	}})
}
