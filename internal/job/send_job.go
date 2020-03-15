package job

import "ramadani.id/jobkue/internal/domain"

type sendJob struct {
	phone, body string
}

func (j *sendJob) Do(message domain.Message, resChan chan<- domain.SendResult) {
	id, err := message.Send(j.phone, j.body)

	resChan <- domain.SendResult{
		ID:  id,
		Err: err,
	}
}

func NewSendJob(phone, body string) domain.SendJob {
	return &sendJob{
		phone: phone,
		body:  body,
	}
}
