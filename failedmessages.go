package websocketchat

type failedmessages interface {
	modifyFailedMessagesAction(action func())
	readFailedMessagesAction(action func())
}
