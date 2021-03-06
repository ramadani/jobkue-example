package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Job interface {
	ID() int
	Do() string
}

type job struct {
	id int
	tm time.Time
}

type req struct {
	ID int `json:"id"`
}

func (j *job) ID() int {
	return j.id
}

func (j *job) Do() string {
	dur := 5
	if j.id%4 == 0 {
		dur = 12
	}

	time.Sleep(time.Duration(dur) * time.Millisecond)
	return fmt.Sprintf("result %d", j.id)
}

func main() {
	pool := make(chan Job, 100)
	result := make(chan string, 100)

	go worker(pool, result)

	log.Println("Serving http")
	http.HandleFunc("/endpoint", wrapper(pool, result))
	http.ListenAndServe(":7000", nil)
}

func worker(jobChan <-chan Job, resChan chan<- string) {
	for job := range jobChan {
		log.Println("doing", job.ID())
		res := job.Do()
		resChan <- res
		log.Println("done", job.ID())
	}
}

func wrapper(jobChan chan<- Job, resChan <-chan string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := req{}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		jobChan <- &job{
			id: req.ID,
			tm: time.Now(),
		}

		resId := <-resChan

		res := make(map[string]string)
		res["id"] = resId
		resp, _ := json.Marshal(res)

		w.Header().Set("Content-Type", "application/json")
		w.Write(resp)
	}
}
