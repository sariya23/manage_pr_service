package pullrequest

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/go-chi/render"
	"github.com/sariya23/manage_pr_service/internal/converters"
	api "github.com/sariya23/manage_pr_service/internal/generated"
	"github.com/sariya23/manage_pr_service/internal/lib/erresponse"
	"github.com/sariya23/manage_pr_service/internal/lib/errorhandler"
	pull_request_validators "github.com/sariya23/manage_pr_service/internal/validators"
)

func (i *PullRequestImplementation) Reassign(w http.ResponseWriter, r *http.Request) {
	const operationPlace = "handlers.pull_request.Reassign"
	log := i.logger.With("operationPlace", operationPlace)
	ctx := r.Context()

	var request api.PostPullRequestReassignJSONRequestBody
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

	if msg, valid := pull_request_validators.ValidatePullRequestReassignRequest(request); !valid {
		log.Warn("invalid request", slog.String("message", msg))
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, erresponse.MakeInvalidResponse(msg))
		return
	}

	pr, newReviewer, err := i.prService.Reassign(ctx, request.PullRequestId, request.OldUserId)
	if status, resp, isError := errorhandler.PullRequestReassign(err); isError {
		w.WriteHeader(status)
		render.JSON(w, r, resp)
		return
	}
	prRes := converters.DomainPullRequestToCreatePullRequestResponse(*pr)
	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, api.PostPullRequestReassign200JSONResponse{
		Pr:         prRes,
		ReplacedBy: newReviewer,
	})
}
