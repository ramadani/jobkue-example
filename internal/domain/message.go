package domain

type Message interface {
	Send(phone, body string) (string, error)
	Delete(id string) error
}
