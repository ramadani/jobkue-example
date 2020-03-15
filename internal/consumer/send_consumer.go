package consumer

import (
	"log"
	"ramadani.id/jobkue/internal/domain"
)

type sendJob struct {
	phone, body string
}

type sendResult struct {
	id  string
	err error
}

type sendConsumer struct {
	msg     domain.Message
	jobChan chan sendJob
	resChan chan sendResult
}

func (j *sendConsumer) Worker() {
	for job := range j.jobChan {
		log.Println("doing", job.phone)
		id, err := j.msg.Send(job.phone, job.body)
		j.resChan <- sendResult{
			id:  id,
			err: err,
		}
		log.Println("done", job.phone, "res", id)
	}
}

func (j *sendConsumer) Queue(phone, body string) bool {
	input := sendJob{phone: phone, body: body}

	select {
	case j.jobChan <- input:
		return true
	default:
		return false
	}
}

func (j *sendConsumer) Result() (string, error) {
	res := <-j.resChan

	return res.id, res.err
}

func NewSendConsumer(msg domain.Message) domain.SendConsumer {
	return &sendConsumer{
		msg:     msg,
		jobChan: make(chan sendJob, 1000),
		resChan: make(chan sendResult, 1000),
	}
}
