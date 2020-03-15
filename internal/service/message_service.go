package service

import "ramadani.id/jobkue/internal/domain"

type messageService struct{}

func (s *messageService) Send(phone, body string) (string, error) {
	id := phone

	return id, nil
}

func (s *messageService) Delete(id string) error {
	return nil
}

func NewMessageService() domain.Message {
	return &messageService{}
}
