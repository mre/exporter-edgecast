package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		resp := Metrics{Namespace: "MY_METRICS", Subsystem: "MY_SUBSYS", Name: "my_own_request_latency", Help: "this is my data"}
		json.NewEncoder(writer).Encode(resp)
	})

	log.Fatal(http.ListenAndServe(":8000", nil))
}

type Metrics struct {
	Namespace string
	Subsystem string
	Name      string
	Help      string
}
