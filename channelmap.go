package websocketchat

type channelmap interface {
	modifyChannelMapAction(action func())
	readChannelMapAction(action func())
}
