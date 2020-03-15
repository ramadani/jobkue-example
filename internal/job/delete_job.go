package job

import "ramadani.id/jobkue/internal/domain"

type deleteJob struct {
	id string
}

func (j *deleteJob) Do(message domain.Message, errChan chan<- error) {
	err := message.Delete(j.id)

	errChan <- err
}

func NewDeleteJob(id string) domain.DeleteJob {
	return &deleteJob{id: id}
}
