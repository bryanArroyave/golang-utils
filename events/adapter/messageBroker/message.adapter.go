package messagebroker

import (
	"encoding/json"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
)

type Message struct {
	message *message.Message
}

func NewBrokerMessage[T any](value T) (*Message, error) {

	vBytes, err := json.Marshal(value)
	if err != nil {
		return nil, err
	}

	return &Message{
		message.NewMessage(watermill.NewUUID(), vBytes),
	}, nil

}

func (m *Message) GetPayload() []byte {
	return m.message.Payload
}

func (m *Message) GetMetadata() map[string]string {
	return m.message.Metadata
}

func (m *Message) GetUUID() string {
	return m.message.UUID
}

func (m *Message) GetMessage() *message.Message {
	return m.message
}

func FromMessage(msg *message.Message) *Message {

	return &Message{
		message: msg,
	}
}
