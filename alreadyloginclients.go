package websocketchat

type alreadyloginclients interface {
	modifyAlreadyLoginClientsAction(action func())
	readAlreadyLoginClientsAction(action func())
	checkUserStatus(username string) bool
}
