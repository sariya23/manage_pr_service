package apiusers

import (
	"errors"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-chi/render"
	"github.com/sariya23/manage_pr_service/internal/converters"
	api "github.com/sariya23/manage_pr_service/internal/generated"
	"github.com/sariya23/manage_pr_service/internal/outerror"
	"github.com/sariya23/manage_pr_service/internal/utils/erresponse"
	validators "github.com/sariya23/manage_pr_service/internal/validators/handlers/users"
)

const (
	URLParamUserUserID = "userID"
)

func (i *UsersImplementation) GetReview(w http.ResponseWriter, r *http.Request, params api.GetUsersGetReviewRequestObject) {
	ctx := r.Context()
	const operationPlace = "handlers.users.GetReview"
	log := i.logger.With("operationPlace", operationPlace)

	if msg, valid := validators.ValidateGetUserReviewRequest(params); !valid {
		log.Warn("invalid request", slog.String("user_id", params.Params.UserId), slog.String("message", msg))
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, erresponse.MakeInvalidResponse(msg))
	}

	userID, _ := strconv.Atoi(params.Params.UserId)
	pullRequests, err := i.userService.GetUserReviews(ctx, int64(userID))
	if err != nil {
		if errors.Is(err, outerror.ErrUserNotFound) {
			log.Warn("user not found", slog.Int("user_id", userID))
			w.WriteHeader(http.StatusNotFound)
			render.JSON(w, r, erresponse.MakeInvalidResponse("user not found"))
			return
		}
		log.Error("unexpected error", slog.Int("user_id", userID), slog.String("error", err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, erresponse.MakeInternalResponse("internal server error"))
		return
	}

	prResponse := converters.MultiDomainPullRequestToGetReviewResponse(pullRequests)
	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, api.GetUsersGetReview200JSONResponse{PullRequests: prResponse, UserId: params.Params.UserId})
}
