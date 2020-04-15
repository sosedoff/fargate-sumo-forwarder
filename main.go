package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func newHandler(authToken string, f forwarder) http.HandlerFunc {
	authToken = "Splunk " + authToken

	return func(rw http.ResponseWriter, r *http.Request) {
		log.Printf("addr=%q host=%q method=%q path=%q agent=%q\n",
			r.RemoteAddr,
			r.Host,
			r.Method,
			r.URL.Path,
			r.UserAgent(),
		)

		if r.Method == http.MethodOptions {
			rw.WriteHeader(200)
			return
		}

		if authToken != "" {
			if r.Header.Get("Authorization") != authToken {
				log.Println("provided auth token is invalid!")

				// Still return 200 so that docker wont retry sending logs
				rw.WriteHeader(200)
				return
			}
		}
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
	if port == "" {
		port = "5000"
	}

	collectorURL := envStringVar("COLLECTOR_URL")
	if collectorURL == "" {
		log.Fatal("COLLECTOR_URL is not set")
	}

	workers := envIntVar("WORKERS")
	if workers == 0 {
		workers = 1
	}

	authToken := envStringVar("AUTH_TOKEN")
	if authToken == "" {
		log.Println("WARNING: AUTH_TOKEN is not set!")
	}

	forwarder := newForwarder(collectorURL, workers)
	go forwarder.start()

	http.HandleFunc("/", newHandler(authToken, forwarder))

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
