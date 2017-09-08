package websocketchat

import "net/http"

type getfriends interface {
	getFriendsHandler(w http.ResponseWriter, r *http.Request)
	serverGetFriends(username string) ([]string, error)
}
