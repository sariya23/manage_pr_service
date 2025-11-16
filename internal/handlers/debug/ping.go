package apidebug

import (
	"net/http"

	"github.com/sariya23/manage_pr_service/internal/lib/request"
)

func (i DebugImplementation) GetDebugPing(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	w.Header().Set("requestID", request.GetIDKey(ctx))
	w.Write([]byte("{\"msg\":\"pong\"}"))
}
