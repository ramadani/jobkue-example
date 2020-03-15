package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Job interface {
	Do(chan<- string)
}

type job struct {
	id int
	tm time.Time
}

type req struct {
	ID int `json:"id"`
}

func (j *job) Do(res chan<- string) {
	log.Println("doing", j.id)
	time.Sleep(100 * time.Millisecond)
	res <- fmt.Sprintf("result %d", j.id)
	log.Println("done", j.id)
}

var pool chan Job
var result chan string

func main() {
	pool = make(chan Job, 10)
	result = make(chan string, 10)

	go worker(pool, result)

	log.Println("Serving http")
	http.HandleFunc("/endpoint", handler)
	http.ListenAndServe(":7000", nil)
}

func worker(jobChan <-chan Job, resChan chan<- string) {
	for job := range jobChan {
		job.Do(resChan)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	req := req{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	go func() {
		pool <- &job{
			id: req.ID,
			tm: time.Now(),
		}
	}()

	resId := <-result

	res := make(map[string]string)
	res["id"] = resId
	resp, _ := json.Marshal(res)

	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}
