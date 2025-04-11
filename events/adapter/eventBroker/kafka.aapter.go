package eventbroker

import (
	"fmt"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
)

// Tipos de configuración de Kafka
type KafkaConfigType string

const (
	PublisherConfig  KafkaConfigType = "Publisher"
	SubscriberConfig KafkaConfigType = "Subscriber"
)

// Estructura de configuración extendida
type KafkaConfig struct {
	Brokers    []string
	Topic      string
	GroupID    string // Requerido solo para el Subscriber
	ConfigType KafkaConfigType
}

// Adaptador Kafka con Zerolog
type KafkaWatermillAdapter struct {
	Publisher  *kafka.Publisher
	Subscriber *kafka.Subscriber
}

// Constructor
func NewKafkaWatermillAdapter(config *KafkaConfig) (*KafkaWatermillAdapter, error) {
	logger := watermill.NewStdLogger(false, false)
	// Crear publisher o subscriber según el tipo
	switch config.ConfigType {
	case PublisherConfig:
		publisher, err := kafka.NewPublisher(kafka.PublisherConfig{
			Brokers:   config.Brokers,
			Marshaler: kafka.DefaultMarshaler{}, // Serialización por defecto
		}, logger)
		if err != nil {
			return nil, fmt.Errorf("error creando publisher Kafka: %w", err)
		}
		return &KafkaWatermillAdapter{Publisher: publisher}, nil

	case SubscriberConfig:
		if config.GroupID == "" {
			return nil, fmt.Errorf("SubscriberConfig requiere un GroupID")
		}
		subscriber, err := kafka.NewSubscriber(kafka.SubscriberConfig{
			Brokers:                config.Brokers,
			Unmarshaler:            kafka.DefaultMarshaler{}, // Deserialización por defecto
			ConsumerGroup:          config.GroupID,
			InitializeTopicDetails: nil, // Usar valores por defecto
		}, logger)
		if err != nil {
			return nil, fmt.Errorf("error creando subscriber Kafka: %w", err)
		}
		return &KafkaWatermillAdapter{Subscriber: subscriber}, nil

	default:
		return nil, fmt.Errorf("tipo de configuración desconocido: %s", config.ConfigType)
	}
}
