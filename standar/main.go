package main

import (
	"fmt"
	"log"
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

func worker1(jobChan <-chan int) {
	for job := range jobChan {
		fmt.Println(job)
	}
}

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	var one, two string

	go func() {
		time.Sleep(400 * time.Millisecond)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(200 * time.Millisecond)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			one = msg1
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			two = msg2
			fmt.Println("received", msg2)
		}
	}

	log.Println("res", one, two)

	//msg := make(chan string)
	//
	//go func() {
	//	log.Println("send")
	//	msg <- "ping"
	//}()
	//
	//log.Println("receive")
	//m := <-msg
	//fmt.Println(m)

	//jobChan := make(chan int, 2)
	//
	//go worker1(jobChan)
	//
	//jobChan <- 1
	//jobChan <- 3
	//jobChan <- 5
	//
	//time.Sleep(2 * time.Second)

	//for i := 1; i <= 2; i++ {
	//	jobChan <- i
	//}
	//
	//for i := 1; i <= 2; i++ {
	//	fmt.Println(<-jobChan)
	//}

	//jobChan <- 1
	//jobChan <- 2
	//
	//fmt.Println(<-jobChan)
	//fmt.Println(<-jobChan)
	//
	//jobChan <- 4
	//fmt.Println(<-jobChan)

	//numJobs := 5
	//jobChanMap := make(map[string]chan Job)
	//resChanMap := make(map[string]chan Result)
	//
	//jobChanMap["test"] = make(chan Job, numJobs)
	//jobChanMap["test2"] = make(chan Job, numJobs)
	//resChanMap["test"] = make(chan Result, numJobs)
	//resChanMap["test2"] = make(chan Result, numJobs)
	//
	//for k, jobChan := range jobChanMap {
	//	resChan := resChanMap[k]
	//
	//	go worker(jobChan, resChan)
	//
	//	for i := 1; i <= numJobs+5; i++ {
	//		jobChan <- Job{
	//			ID:  i,
	//			Key: k,
	//		}
	//	}
	//	close(jobChan)
	//
	//	for i := 1; i <= numJobs+5; i++ {
	//		r := <-resChan
	//		fmt.Println("worker", r.Key, "result", r.ID)
	//	}
	//}
}

func worker(jobChan <-chan Job, resChan chan<- Result) {
	for job := range jobChan {
		fmt.Println("worker", job.Key, "started job", job.ID)
		time.Sleep(50 * time.Millisecond)
		fmt.Println("worker", job.Key, "finished job", job.ID)
		resChan <- Result{ID: job.ID, Key: job.Key}
	}
}
