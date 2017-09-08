package websocketchat

import "net/http"

type addfriends interface {
	addFriendHandler(w http.ResponseWriter, r *http.Request)
}
