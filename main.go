package websocketchat

import "net/http"

type Message struct {
}

type mainHandler interface {
	handler(w http.ResponseWriter, r *http.Request)
	indexFileHandler(w http.ResponseWriter, r *http.Request)
	boardcaseMessage()
	exitFromChannel(message Message)
	sendToChannel(message Message)
	pushChannelMessage(message Message, channelName string)
	sendMessage(message Message)
	pushFaileMessage(user string, message Message)
}
