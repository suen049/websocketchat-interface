package websocketchat

import "net/http"

type register interface {
	registerHandler(w http.ResponseWriter, r *http.Request)
}
