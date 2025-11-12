package apidebug

import "net/http"

func (i *DebugImplementation) Ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
	return
}
