package main

import (
	"fmt"
	"time"
)

type Job struct {
	ID  int
	Key string
}

type Result struct {
	ID  int
	Key string
}

func main() {
	numJobs := 5
	jobChanMap := make(map[string]chan Job)
	resChanMap := make(map[string]chan Result)

	jobChanMap["test"] = make(chan Job, numJobs)
	jobChanMap["test2"] = make(chan Job, numJobs)
	resChanMap["test"] = make(chan Result, numJobs)
	resChanMap["test2"] = make(chan Result, numJobs)

	for k, jobChan := range jobChanMap {
		resChan := resChanMap[k]

		go worker(jobChan, resChan)

		for i := 1; i <= numJobs; i++ {
			jobChan <- Job{
				ID:  i,
				Key: k,
			}
		}
		close(jobChan)

		for i := 1; i <= numJobs; i++ {
			r := <-resChan
			fmt.Println("worker", r.Key, "result", r.ID)
		}
	}
}

func worker(jobChan <-chan Job, resChan chan<- Result) {
	for job := range jobChan {
		fmt.Println("worker", job.Key, "started job", job.ID)
		time.Sleep(50 * time.Millisecond)
		fmt.Println("worker", job.Key, "finished job", job.ID)
		resChan <- Result{ID: job.ID, Key: job.Key}
	}
}
