package service

import (
	"ramadani.id/jobkue/internal/domain"
	"strconv"
	"time"
)

type messageService struct{}

func (s *messageService) Send(phone, body string) (string, error) {
	id := phone
	lastStr := phone[len(phone)-1:]
	last, _ := strconv.Atoi(lastStr)
	dur := time.Duration(1)
	if last%4 == 0 {
		dur = time.Duration(3)
	}

	time.Sleep(dur * time.Millisecond)
	return id, nil
}

func (s *messageService) Delete(id string) error {
	return nil
}

func NewMessageService() domain.Message {
	return &messageService{}
}
