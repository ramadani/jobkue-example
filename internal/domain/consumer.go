package domain

type SendConsumer interface {
	Queue(phone, body string)
	Result() (string, error)
	Worker()
}
