package apidebug

import "net/http"

func (i *DebugImplementation) Ping(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("pong"))
}
