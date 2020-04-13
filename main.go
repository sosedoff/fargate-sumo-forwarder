package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func newHandler(f forwarder) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		defer rw.WriteHeader(200)

		startTime := time.Now()
		log.Println(r.Method, r.URL.String(), time.Since(startTime).String())

		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("body read error:", err)
			return
		}
		if len(data) == 0 {
			return
		}

		f.queue(data)
	}
}

func envStringVar(key string) string {
	return os.Getenv(key)
}

func envIntVar(key string) int {
	var val int
	fmt.Sscanf(os.Getenv(key), "%d", &val)
	return val
}

func main() {
	port := envStringVar("PORT")
	collectorURL := envStringVar("COLLECTOR_URL")
	workers := envIntVar("COLLECTOR_WORKERS")

	forwarder := newForwarder(collectorURL, workers)
	go forwarder.start()

	http.HandleFunc("/", newHandler(forwarder))

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
