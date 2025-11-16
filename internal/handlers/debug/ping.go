package apidebug

import (
	"net/http"

	"github.com/sariya23/manage_pr_service/internal/middleware"
)

func (i DebugImplementation) GetDebugPing(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	requestID, ok := ctx.Value(middleware.RequestIDKey).(string)
	if !ok {
		requestID = "-1"
	}
	w.Header().Set("requestID", requestID)
	w.Write([]byte("{\"msg\":\"pong\"}"))
}
