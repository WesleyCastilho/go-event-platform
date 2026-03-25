package port

type Publisher interface {
	Publish(queue string, body []byte) error
}
