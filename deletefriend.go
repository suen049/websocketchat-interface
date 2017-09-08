package websocketchat

import "net/http"

type deletefriend interface {
	deleteFriendHandler(w http.ResponseWriter, r *http.Request)
}
