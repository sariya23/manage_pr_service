package apidebug

import "net/http"

func (i DebugImplementation) GetDebugPing(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("{\"msg\":\"pong\"}"))
}
