package pullrequest

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"github.com/go-chi/render"
	"github.com/sariya23/manage_pr_service/internal/converters"
	api "github.com/sariya23/manage_pr_service/internal/generated"
	"github.com/sariya23/manage_pr_service/internal/lib/erresponse"
	"github.com/sariya23/manage_pr_service/internal/outerror"
)

func (i *PullRequestImplementation) Merge(w http.ResponseWriter, r *http.Request) {
	const operationPlace = "handlers.pull_request.Merge"
	log := i.logger.With("operationPlace", operationPlace)
	ctx := r.Context()

	var request api.PostPullRequestMergeJSONRequestBody
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

	if request.PullRequestId == "" {
		log.Warn("invalid request", slog.String("message", "pull_request_id is required"))
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, erresponse.MakeInvalidResponse("pull_request_id is required"))
		return
	}

	pr, err := i.prService.Merge(ctx, request.PullRequestId)
	if err != nil {
		if errors.Is(err, outerror.ErrPullRequestNotFound) {
			w.WriteHeader(http.StatusNotFound)
			render.JSON(w, r, erresponse.MakeNotFoundResponse("resource not found"))
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, erresponse.MakeInternalResponse("internal server error"))
		return
	}
	prRes := converters.DomainPullRequestToCreatePullRequestResponse(*pr)
	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, api.PostPullRequestMerge200JSONResponse{
		Pr: &prRes,
	})
}
