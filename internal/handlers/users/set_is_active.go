package apiusers

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"github.com/go-chi/render"
	"github.com/sariya23/manage_pr_service/internal/converters"
	api "github.com/sariya23/manage_pr_service/internal/generated"
	"github.com/sariya23/manage_pr_service/internal/outerror"
	"github.com/sariya23/manage_pr_service/internal/utils/erresponse"
	validators "github.com/sariya23/manage_pr_service/internal/validators/handlers/users"
)

func (i *UsersImplementation) SetIsActive(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	const operationPlace = "handlers.users.SetIsActive"
	log := i.logger.With("operationPlace", operationPlace)
	var request api.PostUsersSetIsActiveJSONRequestBody

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

	if msg, valid := validators.ValidateSetIsActiveUserRequest(request); !valid {
		log.Warn("invalid request", slog.String("user_id", request.UserId), slog.String("message", msg))
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, erresponse.MakeInvalidResponse(msg))
		return
	}
	domainUser, err := i.userService.SetIsActive(ctx, request.UserId, request.IsActive)
	if err != nil {
		if errors.Is(err, outerror.ErrUserNotFound) {
			log.Warn("user not found", slog.String("user_id", request.UserId))
			w.WriteHeader(http.StatusNotFound)
			render.JSON(w, r, erresponse.MakeNotFoundResponse("user not found"))
			return
		}
		log.Error("unexpected error", slog.String("user_id", request.UserId), slog.String("error", err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, erresponse.MakeInternalResponse("internal server error"))
		return
	}

	responseUser := converters.DomainUserToIsActiveResponseUser(domainUser)
	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, api.PostUsersSetIsActive200JSONResponse{User: &responseUser})
	return
}
