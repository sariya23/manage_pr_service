package apiteams

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"github.com/go-chi/render"
	api "github.com/sariya23/manage_pr_service/internal/generated"
	"github.com/sariya23/manage_pr_service/internal/lib/erresponse"
	"github.com/sariya23/manage_pr_service/internal/lib/request"
	"github.com/sariya23/manage_pr_service/internal/outerror"
)

func (i TeamsImplementation) PostTeamDeactivate(w http.ResponseWriter, r *http.Request) {
	const operationPlace = "handlers.pull_request.PostTeamDeactivate"
	log := i.logger.With("operationPlace", operationPlace)
	ctx := r.Context()
	requestID := request.GetIDKey(ctx)
	log = log.With("request_id", requestID)
	w.Header().Set("requestID", requestID)

	var rq api.PostTeamDeactivateJSONRequestBody
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

	if rq.TeamName == "" || len(rq.UserIds) == 0 {
		log.Warn("no team name or user ids provided")
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, erresponse.MakeInvalidResponse("no team name or user ids provided"))
		return
	}

	err := i.teamsService.Deactivate(ctx, rq.TeamName, rq.UserIds)
	if err != nil {
		if errors.Is(err, outerror.ErrTeamNotFound) {
			w.WriteHeader(http.StatusNotFound)
			render.JSON(w, r, erresponse.MakeNotFoundResponse("team not found"))
			return
		} else if errors.Is(err, outerror.ErrUserNotInTeam) {
			w.WriteHeader(http.StatusBadRequest)
			render.JSON(w, r, erresponse.MakeInvalidResponse("one of user not in team"))
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		log.Error("error deactivating team", slog.String("error", err.Error()))
		render.JSON(w, r, erresponse.MakeInternalResponse("error deactivating team"))
		return
	}

	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, api.PostTeamDeactivate200Response{})
}
