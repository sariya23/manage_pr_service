package apiusers

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/go-chi/render"
	"github.com/sariya23/manage_pr_service/internal/converters"
	api "github.com/sariya23/manage_pr_service/internal/generated"
	"github.com/sariya23/manage_pr_service/internal/lib/erresponse"
	"github.com/sariya23/manage_pr_service/internal/outerror"

	validators "github.com/sariya23/manage_pr_service/internal/validators/handlers/users"
)

func (i *UsersImplementation) GetReview(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	const operationPlace = "handlers.users.GetReview"
	log := i.logger.With("operationPlace", operationPlace)
	query := r.URL.Query()
	userID := query.Get("user_id")
	params := api.GetUsersGetReviewRequestObject{}
	params.Params = api.GetUsersGetReviewParams{UserId: userID}
	if msg, valid := validators.ValidateGetUserReviewRequest(params); !valid {
		log.Warn("invalid request", slog.String("user_id", params.Params.UserId), slog.String("message", msg))
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, erresponse.MakeInvalidResponse(msg))
	}

	pullRequests, err := i.userService.GetReviews(ctx, params.Params.UserId)
	if err != nil {
		if errors.Is(err, outerror.ErrUserNotFound) {
			log.Warn("user not found", slog.String("user_id", params.Params.UserId))
			w.WriteHeader(http.StatusNotFound)
			render.JSON(w, r, erresponse.MakeNotFoundResponse("user not found"))
			return
		}
		log.Error("unexpected error", slog.String("user_id", params.Params.UserId), slog.String("error", err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, erresponse.MakeInternalResponse("internal server error"))
		return
	}

	prResponse := converters.MultiDomainPullRequestToGetReviewResponse(pullRequests)
	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, api.GetUsersGetReview200JSONResponse{PullRequests: prResponse, UserId: params.Params.UserId})
}
