package api_pull_requests

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/go-chi/render"
	"github.com/sariya23/manage_pr_service/internal/converters"
	api "github.com/sariya23/manage_pr_service/internal/generated"
	"github.com/sariya23/manage_pr_service/internal/lib/erresponse"
	"github.com/sariya23/manage_pr_service/internal/lib/errorhandler"
	"github.com/sariya23/manage_pr_service/internal/models/domain"
	"github.com/sariya23/manage_pr_service/internal/models/dto"
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

	pullRequest, reviewers, err := i.prService.CreatePullRequestAndAssignReviewers(ctx, dto.FromCreatePullRequestHTTP(request))
	if status, resp, isError := errorhandler.PullRequestCreate(err); isError {
		w.WriteHeader(status)
		render.JSON(w, r, resp)
		return
	}
	prRes := converters.DomainPullRequestToCreatePullRequestResponse(*pullRequest)
	prRes.AssignedReviewers = domain.UserIDs(reviewers)
	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, api.PostPullRequestMerge200JSONResponse{
		Pr: &prRes,
	})
}
