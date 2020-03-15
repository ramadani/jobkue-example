package main

import (
	"encoding/json"
	"log"
	"net/http"
	"ramadani.id/jobkue/internal/consumer"
	"ramadani.id/jobkue/internal/domain"
	"ramadani.id/jobkue/internal/service"
)

type SendReq struct {
	Phone string `json:"phone"`
	Body  string `json:"body"`
}

type SendResp struct {
	ID string `json:"id"`
}

type handler struct {
	msg domain.Message
}

func main() {
	msg := service.NewMessageService()
	sendConsWorker := consumer.NewSendConsumerWorker(msg)
	msg = service.NewMessageServiceJob(sendConsWorker)
	workers := []domain.Worker{
		sendConsWorker,
	}

	for _, worker := range workers {
		go worker.Worker()
	}

	h := &handler{msg}

	log.Println("Starting app")
	http.HandleFunc("/send", h.Send)
	http.ListenAndServe(":5000", nil)
}

func (h *handler) Send(w http.ResponseWriter, r *http.Request) {
	req := &SendReq{}

	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := h.msg.Send(req.Phone, req.Body)
	if err != nil && err == domain.OverCapsErr {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(&SendResp{ID: id})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
