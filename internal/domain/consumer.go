package domain

import "errors"

var OverCapsErr = errors.New("over capacity")

type SendConsumer interface {
	Queue(phone, body string) bool
	Result() (string, error)
}
