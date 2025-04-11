package eventbroker

import (
	"fmt"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-amqp/v3/pkg/amqp"
)

// vist: https://watermill.io/pubsubs/amqp/
type AMQPConfigType string

const (
	DurablePubSubConfig    AMQPConfigType = "DurablePubSub"
	NonDurablePubSubConfig AMQPConfigType = "NonDurablePubSub"
	DurableQueueConfig     AMQPConfigType = "DurableQueue"
	NonDurableQueueConfig  AMQPConfigType = "NonDurableQueue"
	DurableTopicConfig     AMQPConfigType = "DurableTopic"
	NonDurableTopicConfig  AMQPConfigType = "NonDurableTopic"
)

type AMQPConfig struct {
	URI        string
	Exchange   string
	Queue      string
	RoutingKey string
	ConfigType AMQPConfigType
}

type AMQPWatermillAdapter struct {
	Publisher  *amqp.Publisher
	Subscriber *amqp.Subscriber
}

func NewAMQPWatermillAdapter(config *AMQPConfig) (*AMQPWatermillAdapter, error) {

	logger := watermill.NewStdLogger(false, false)
	cfg, err := getConfig(config)
	if err != nil {
		return nil, err
	}

	publisher, err := amqp.NewPublisher(*cfg, logger)
	if err != nil {
		return nil, fmt.Errorf("error creando publisher: %w", err)
	}

	subscriber, err := amqp.NewSubscriber(*cfg, logger)
	if err != nil {
		return nil, fmt.Errorf("error creando subscriber: %w", err)
	}

	return &AMQPWatermillAdapter{
		Publisher:  publisher,
		Subscriber: subscriber,
	}, nil
}

func getConfig(config *AMQPConfig) (*amqp.Config, error) {
	switch config.ConfigType {
	case DurablePubSubConfig:
		cfg := amqp.NewDurablePubSubConfig(config.URI, amqp.GenerateQueueNameTopicName)
		return &cfg, nil
	case NonDurablePubSubConfig:
		cfg := amqp.NewNonDurablePubSubConfig(config.URI, amqp.GenerateQueueNameTopicName)
		return &cfg, nil
	case DurableQueueConfig:
		cfg := amqp.NewDurableQueueConfig(config.URI)
		return &cfg, nil
	case NonDurableQueueConfig:
		cfg := amqp.NewNonDurableQueueConfig(config.URI)
		return &cfg, nil
	case DurableTopicConfig:
		if config.Exchange == "" {
			return nil, fmt.Errorf("DurableTopicConfig requiere un exchange")
		}
		cfg := amqp.NewDurableTopicConfig(config.URI, config.Exchange, config.RoutingKey)
		return &cfg, nil
	case NonDurableTopicConfig:
		if config.Exchange == "" {
			return nil, fmt.Errorf("NonDurableTopicConfig requiere un exchange")
		}
		cfg := amqp.NewNonDurableTopicConfig(config.URI, config.Exchange, config.RoutingKey)
		return &cfg, nil
	default:
		return nil, fmt.Errorf("tipo de configuración desconocido: %s", config.ConfigType)
	}
}
