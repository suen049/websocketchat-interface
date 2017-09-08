package websocketchat

import "net/http"

type logout interface {
	logoutHandler(w http.ResponseWriter, r *http.Request)
}
