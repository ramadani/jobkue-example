package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	log.Println("Calling http")

	for i := 1; i <= 10; i++ {
		body, err := json.Marshal(map[string]int{
			"id": i,
		})
		if err != nil {
			log.Fatal(err)
		}

		resp, err := http.Post("http://localhost:7000/endpoint", "application/json", bytes.NewBuffer(body))
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		bytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		data := make(map[string]interface{})
		err = json.Unmarshal(bytes, &data)
		if err != nil {
			log.Fatal(err)
		}

		log.Println(data)
	}
}
