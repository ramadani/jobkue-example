package main

import (
	"encoding/json"
	"fmt"
	vegeta "github.com/tsenart/vegeta/lib"
	"log"
	"net/http"
	"strings"
	"time"
)

type req struct {
	ID int `json:"id"`
}

func main() {
	attack()
}

func attack() {
	start := time.Now()
	rate := vegeta.Rate{Freq: 300, Per: time.Second}
	duration := 5 * time.Second

	targets := make([]vegeta.Target, 0)
	for i := 0; i < 300; i++ {
		header := make(http.Header)
		header.Set("Content-Type", "application/json")
		data := &req{
			i + 1,
		}
		body, _ := json.Marshal(data)

		targets = append(targets, vegeta.Target{
			Method: "POST",
			URL:    "http://localhost:7000/endpoint",
			Header: header,
			Body:   body,
		})
	}

	targeter := vegeta.NewStaticTargeter(targets...)
	attacker := vegeta.NewAttacker()

	var metrics vegeta.Metrics
	for res := range attacker.Attack(targeter, rate, duration, "Attack") {
		metrics.Add(res)
	}
	metrics.Close()

	getMetric(metrics)

	end := time.Now()
	dif := end.Sub(start)
	log.Println("total duration", dif)
}

func getMetric(metrics vegeta.Metrics) {
	statusCodes := make([]string, len(metrics.StatusCodes))
	for k, v := range metrics.StatusCodes {
		statusCodes = append(statusCodes, fmt.Sprintf("%s:%d", k, v))
	}

	fmt.Printf("Requests [total, rate, throughput]  %d %.2f %.2f\n", metrics.Requests, metrics.Rate, metrics.Throughput)
	fmt.Printf("Durations [total, attack, wait]  %s %s %s\n", metrics.Duration,
		metrics.Latencies.Total, metrics.Wait)
	fmt.Printf("Latencies [mean, 50, 95, 99, max]  %s %s %s %s %s\n", metrics.Latencies.Mean,
		metrics.Latencies.P50, metrics.Latencies.P95, metrics.Latencies.P99, metrics.Latencies.Max)
	fmt.Printf("Bytes In [total, mean]  %d %.2f\n", metrics.BytesIn.Total, metrics.BytesIn.Mean)
	fmt.Printf("Bytes Out [total, mean]  %d %.2f\n", metrics.BytesOut.Total, metrics.BytesOut.Mean)
	fmt.Printf("Success [ratio]  %.2f percent\n", metrics.Success*100)
	fmt.Printf("Status Codes [code:count]  %s\n", strings.Join(statusCodes, " "))

	if len(metrics.Errors) > 0 {
		fmt.Printf("%s\n", strings.Join(metrics.Errors, "\n"))
	}
}
