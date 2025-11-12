package apiusers

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-chi/render"
	"github.com/sariya23/manage_pr_service/internal/converters"
	api "github.com/sariya23/manage_pr_service/internal/generated"
	"github.com/sariya23/manage_pr_service/internal/handlers"
	"github.com/sariya23/manage_pr_service/internal/outerror"
)

func (i *UsersImplementation) SetIsActive(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	const operationPlace = "handlers.users.SetIsActive"
	log := i.logger.With("operationPlace", operationPlace)
	var request api.PostUsersSetIsActiveJSONRequestBody

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		log.Error("error decoding request body", slog.String("error", err.Error()))
		errorResp := api.ErrorResponse{}
		errorResp.Error.Code = handlers.NOT_FOUND
		errorResp.Error.Message = "user not found"
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, errorResp)
		return
	}
	defer func() {
		if err := r.Body.Close(); err != nil {
			log.Warn("error closing body", slog.String("error", err.Error()))
		}
	}()

	if request.UserId == "" {
		log.Warn("empty user_id")
		errorResp := api.ErrorResponse{}
		errorResp.Error.Code = handlers.MISSING_REQUIRED_FIELD
		errorResp.Error.Message = "user_id is required"
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, errorResp)
		return
	}

	userID, err := strconv.Atoi(request.UserId)
	if err != nil {
		log.Warn("user_id is not an integer",
			slog.String("user_id", request.UserId),
			slog.String("error", err.Error()))
		errorResp := api.ErrorResponse{}
		errorResp.Error.Code = handlers.TYPE_MISMATCH
		errorResp.Error.Message = "user_id must be an integer"
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, errorResp)
		return
	}

	domainUser, err := i.userService.SetIsActive(ctx, int64(userID), request.IsActive)
	if err != nil {
		if errors.Is(err, outerror.ErrUserNotFound) {
			log.Warn("user not found", slog.String("user_id", request.UserId))
			errorResp := api.ErrorResponse{}
			errorResp.Error.Code = handlers.NOT_FOUND
			errorResp.Error.Message = "user not found"
			w.WriteHeader(http.StatusNotFound)
			render.JSON(w, r, errorResp)
			return
		}
		log.Error("unexpected error", slog.String("user_id", request.UserId), slog.String("error", err.Error()))
		errorResp := api.ErrorResponse{}
		errorResp.Error.Code = handlers.INTERNAL
		errorResp.Error.Message = "internal error"
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, errorResp)
		return
	}

	responseUser := converters.DomainUserToIsActiveResponseUser(domainUser)
	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, responseUser)
	return
}
