package apiteams

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/go-chi/render"
	api "github.com/sariya23/manage_pr_service/internal/generated"
	"github.com/sariya23/manage_pr_service/internal/utils/erresponse"
)

func (i *TeamsImplementation) Add(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	const operationPlace = "handlers.teams.Add"
	log := i.logger.With("operationPlace", operationPlace)

	var request api.PostTeamAddJSONRequestBody
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

	// team_name = ""
	// validate user id
	// validate username
	// validate is_active
}
