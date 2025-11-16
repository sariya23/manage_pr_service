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
	"github.com/sariya23/manage_pr_service/internal/lib/request"
	teamsvalidators "github.com/sariya23/manage_pr_service/internal/validators"
)

func (i TeamsImplementation) PostTeamAdd(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	const operationPlace = "handlers.teams.Add"
	log := i.logger.With("operationPlace", operationPlace)
	requestID := request.GetIDKey(ctx)
	log = log.With("request_id", requestID)
	w.Header().Set("requestID", requestID)

	var rq api.PostTeamAddJSONRequestBody
	if err := json.NewDecoder(r.Body).Decode(&rq); err != nil {
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

	if msg, valid := teamsvalidators.ValidateTeamAddRequest(rq); !valid {
		log.Warn("invalid request", slog.String("message", msg))
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, erresponse.MakeInvalidResponse(msg))
		return
	}

	members, err := i.teamsService.Add(ctx, rq.TeamName, converters.MultiAddTeamUserToDomainUser(rq.Members))
	if status, resp, isError := errorhandler.TeamAdd(err, rq.TeamName); isError {
		w.WriteHeader(status)
		render.JSON(w, r, resp)
		return
	}
	log.Info("team created", slog.String("team_name", rq.TeamName))
	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, api.PostTeamAdd201JSONResponse{Team: &api.Team{
		Members:  converters.MultiDomainUserToAddTeamResponse(members),
		TeamName: rq.TeamName,
	}})
}
