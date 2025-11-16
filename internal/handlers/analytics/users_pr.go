package analytics

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/render"
	api "github.com/sariya23/manage_pr_service/internal/generated"
)

func (i *AnalyticsImplementation) UsersPRs(w http.ResponseWriter, r *http.Request) {
	const operationPlace = "handlers.analytics.UsersPRs"
	log := i.log.With("operationPlace", operationPlace)
	ctx := r.Context()

	m, err := i.PullRequestService.GroupPullRequestsByAssignedReviewer(ctx)
	if err != nil {
		log.Error("could not get group prs by assigned reviewer", slog.String("error", err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, api.GetAnalyticsUsersPRs200JSONResponse(m))
}
