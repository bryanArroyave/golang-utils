package factory

import (
	"fmt"

	"github.com/ThreeDotsLabs/watermill/message"
	adapter "github.com/bryanArroyave/golang-utils/events/adapter/eventBroker"
	"github.com/bryanArroyave/golang-utils/events/enums"
)

type Config struct {
	BrokerType enums.BrokerType

	// Kafka Config
	Brokers []string

	// RabbitMQ Config
	URI string

	// Channels Config (for in-memory messaging)
	BufferSize int
}

type MessageBroker struct {
	Publisher  message.Publisher
	Subscriber message.Subscriber
}

type FactoryConfig struct {
	AMQP      *adapter.AMQPConfig
	Kafka     *adapter.KafkaConfig
	GoChannel *adapter.GoChannelConfig
}

func NewMessageBroker(adapterType enums.BrokerType, config *FactoryConfig) (*MessageBroker, error) {

	switch string(adapterType) {
	case string(enums.RabbitMQ):
		if config.AMQP == nil {
			return nil, fmt.Errorf("configuracion de RabbitMQ no proporcionada")
		}
		amqpadapter, err := adapter.NewAMQPWatermillAdapter(config.AMQP)

		if err != nil {
			return nil, fmt.Errorf("error creando adaptador RabbitMQ: %w", err)
		}

		return &MessageBroker{
			Publisher:  amqpadapter.Publisher,
			Subscriber: amqpadapter.Subscriber,
		}, nil

	case string(enums.Kafka):
		if config.Kafka == nil {
			return nil, fmt.Errorf("configuracion de Kafka no proporcionada")
		}

		kafkaadapter, err := adapter.NewKafkaWatermillAdapter(config.Kafka)
		if err != nil {
			return nil, fmt.Errorf("error creando adaptador Kafka: %w", err)
		}
		return &MessageBroker{
			Publisher:  kafkaadapter.Publisher,
			Subscriber: kafkaadapter.Subscriber,
		}, nil

	case string(enums.Channels):
		if config.GoChannel == nil {
			return nil, fmt.Errorf("configuracion de GoChannel no proporcionada")
		}

		goadapter := adapter.NewGoChannelWatermillAdapter(config.GoChannel)
		return &MessageBroker{
			Publisher:  goadapter.Publisher,
			Subscriber: goadapter.Subscriber,
		}, nil

	default:
		return nil, fmt.Errorf("tipo de adaptador no soportado: %s", adapterType)
	}
}
