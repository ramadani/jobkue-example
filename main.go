package main

import (
	"log"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {

	jobChanMap := make(map[string]chan Job)
	jobChanMap["test"] = make(chan Job, 5)
	jobChanMap["test2"] = make(chan Job, 5)

	for k, jobChan := range jobChanMap {
		wg.Add(1)
		go worker(jobChan)

		for i := 0; i < 10; i++ {
			if !TryEnqueue(Job{ID: i + 1, Key: k}, jobChan) {
				log.Println("max caps")
			}
		}

		close(jobChan)

		wg.Wait()
	}

	//jobChan := jobChanMap["test"]

	//time.Sleep(20 * time.Second)
}

type Job struct {
	ID  int
	Key string
}

func worker(jobChan <-chan Job) {
	defer wg.Done()

	for job := range jobChan {
		log.Println("start job", job.ID, "key", job.Key)
		time.Sleep(100 * time.Millisecond)
		log.Println("end job", job.ID, "key", job.Key)
	}
}

func TryEnqueue(job Job, jobChan chan<- Job) bool {
	select {
	case jobChan <- job:
		return true
	default:
		return false
	}
}
