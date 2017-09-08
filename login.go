package websocketchat

import "net/http"

type login interface {
	loginHandler(w http.ResponseWriter, r *http.Request)
}
