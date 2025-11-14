package api_pull_requests

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"github.com/go-chi/render"
	"github.com/sariya23/manage_pr_service/internal/converters"
	api "github.com/sariya23/manage_pr_service/internal/generated"
	"github.com/sariya23/manage_pr_service/internal/lib/erresponse"
	"github.com/sariya23/manage_pr_service/internal/models/domain"
	"github.com/sariya23/manage_pr_service/internal/models/dto"
	"github.com/sariya23/manage_pr_service/internal/outerror"
	pull_request_validators "github.com/sariya23/manage_pr_service/internal/validators/handlers/pull_request"
)

func (i *PullRequestImplementation) Create(w http.ResponseWriter, r *http.Request) {
	const operationPlace = "handlers.pull_request.Create"
	log := i.logger.With("operationPlace", operationPlace)
	ctx := r.Context()

	var request api.PostPullRequestCreateJSONRequestBody
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

	if msg, valid := pull_request_validators.ValidatePullRequestCreateRequest(request); !valid {
		log.Warn("invalid request", slog.String("message", msg))
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, erresponse.MakeInvalidResponse(msg))
		return
	}

	pullRequest, reviewers, err := i.prService.CreatePullRequest(ctx, dto.FromCreatePullRequestHTTP(request))
	if err != nil {
		if errors.Is(err, outerror.ErrPullRequestAlreadyExists) {
			w.WriteHeader(http.StatusConflict)
			render.JSON(w, r, erresponse.MakePullRequestAlreadyExistsResponse("PR id already exists"))
			return
		} else if errors.Is(err, outerror.ErrUserNotFound) {
			w.WriteHeader(http.StatusConflict)
			render.JSON(w, r, erresponse.MakeNotFoundResponse("author_id not found"))
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, erresponse.MakeInternalResponse("internal server error"))
		return
	}

	prRes := converters.DomainPullRequestToCreatePullRequestResponse(pullRequest)
	prRes.AssignedReviewers = domain.UserIDs(reviewers)
	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, api.PostPullRequestMerge200JSONResponse{
		Pr: &prRes,
	})
}
