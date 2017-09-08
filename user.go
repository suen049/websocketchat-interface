package websocketchat

type UserInformation struct {
}

type Users struct {
}

type user interface {
	modifyUsersAction(action func())
	readUsersAction(action func())
	checkUserLogin(id string, password string) error
	registerUser(id string, password string) error
	addFriend(id string, friendID string) error
	deleteFriend(id string, friendID string) error
}
