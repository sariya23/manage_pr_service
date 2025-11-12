package apiusers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/sariya23/manage_pr_service/internal/utils/erresponse"
)

const (
	URLParamUserUserID = "userID"
)

func (i *UsersImplementation) GetReview(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	const operationPlace = "handlers.users.GetReview"
	log := i.logger.With("operationPlace", operationPlace)
	userID := chi.URLParam(r, URLParamUserUserID)
	if userID == "" {
		log.Warn("userID parameter is required")
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, erresponse.MakeInvalidResponse("missing query parameter user_id"))
		return
	}

}
