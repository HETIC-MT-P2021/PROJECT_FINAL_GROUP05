package rabbit

type RabbitRepository interface {
	Publish(string) error
	Consume()
}