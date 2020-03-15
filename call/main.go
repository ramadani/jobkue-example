package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func main() {
	log.Println("Calling http")

	for i := 1; i <= 200; i++ {
		body, err := json.Marshal(map[string]interface{}{
			"phone": strconv.Itoa(i),
			"body":  "lorem ipsum",
		})
		if err != nil {
			log.Fatal(err)
		}

		resp, err := http.Post("http://localhost:5000/send", "application/json", bytes.NewBuffer(body))
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
