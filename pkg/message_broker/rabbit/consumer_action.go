package rabbit

type ConsumerAction interface {
	Execute() error
	SetBody(body []byte)
}