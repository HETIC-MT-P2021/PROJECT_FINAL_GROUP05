package rabbit

type ConsumerAction interface {
	Execute([]byte) error
}