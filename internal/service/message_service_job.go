package service

import (
	"ramadani.id/jobkue/internal/domain"
)

type messageServiceJob struct {
	sendConsumer domain.SendConsumer
}

func (s *messageServiceJob) Send(phone, body string) (string, error) {
	if next := s.sendConsumer.Queue(phone, body); !next {
		return "", domain.OverCapsErr
	}

	return s.sendConsumer.Result()
}

func (s *messageServiceJob) Delete(id string) error {
	return nil
}

func NewMessageServiceJob(sendConsumer domain.SendConsumer) domain.Message {
	return &messageServiceJob{sendConsumer: sendConsumer}
}
