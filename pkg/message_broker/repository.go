package rabbit

type MessageBrokerRepository interface {
	Publish(string) error
	Consume(func())
}