package enums

type BrokerType string

const (
	Kafka    BrokerType = "kafka"
	RabbitMQ BrokerType = "rabbitmq"
	Channels BrokerType = "channels"
)
