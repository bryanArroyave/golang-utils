package eventbroker

import (
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
)

type GoChannelConfig struct {
	BufferSize int64
}

// Adaptador Go Channels con Zerolog
type GoChannelWatermillAdapter struct {
	Publisher  *gochannel.GoChannel
	Subscriber *gochannel.GoChannel
}

// Constructor
func NewGoChannelWatermillAdapter(config *GoChannelConfig) *GoChannelWatermillAdapter {
	logger := watermill.NewStdLogger(false, false)
	channel := gochannel.NewGoChannel(gochannel.Config{
		OutputChannelBuffer:            config.BufferSize,
		Persistent:                     true,
		BlockPublishUntilSubscriberAck: true,
	}, logger)

	return &GoChannelWatermillAdapter{
		Publisher:  channel,
		Subscriber: channel,
	}
}
