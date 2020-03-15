package domain

type SendResult struct {
	ID  string
	Err error
}

type SendJob interface {
	Do(message Message, resChan chan<- SendResult)
}

type DeleteJob interface {
	Do(message Message, errChan chan<- error)
}
