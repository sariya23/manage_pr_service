package apiusers

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
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
	userID := chi.URLParam(r, "user_id")
	if msg, valid := validators.ValidateGetUserReviewRequest(userID); !valid {
		log.Warn("invalid request", slog.String("user_id", userID), slog.String("message", msg))
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, erresponse.MakeInvalidResponse(msg))
		return
	}

	pullRequests, err := i.userService.GetReviews(ctx, userID)
	if err != nil {
		if errors.Is(err, outerror.ErrUserNotFound) {
			log.Warn("user not found", slog.String("user_id", userID))
			w.WriteHeader(http.StatusNotFound)
			render.JSON(w, r, erresponse.MakeNotFoundResponse("user not found"))
			return
		}
		log.Error("unexpected error", slog.String("user_id", userID), slog.String("error", err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, erresponse.MakeInternalResponse("internal server error"))
		return
	}

	prResponse := converters.MultiDomainPullRequestToGetReviewResponse(pullRequests)
	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, api.GetUsersGetReview200JSONResponse{PullRequests: prResponse, UserId: userID})
}
