package main

import (
	"log"
	"net/http"
	"sync"
)

type forwarder struct {
	collector string
	client    *http.Client
	data      chan []byte
	workers   int
}

func newForwarder(collector string, workers int) forwarder {
	return forwarder{
		collector: collector,
		client:    http.DefaultClient,
		data:      make(chan []byte, 1024),
		workers:   workers,
	}
}

func (f forwarder) queue(data []byte) {
	f.data <- data
}

func (f forwarder) work(id int, wg *sync.WaitGroup) {
	log.Println("starting worker", id)
	defer func() {
		log.Println("stopping worker", id)
		defer wg.Done()
	}()

	for data := range f.data {
		log.Printf("processing: %s\n", data)

		payloads := parsePayloads(data)
		if len(payloads) == 0 {
			continue
		}

		sumo := sumopayload{
			lines: make([]string, len(payloads)),
		}

		for idx, p := range payloads {
			sumo.lines[idx] = p.Event.Line
			sumo.host = p.Host
			sumo.sourceName = p.Source
			sumo.sourceCategory = p.SourceType
		}

		go func() {
			if err := sumo.send(f.client, f.collector); err != nil {
				log.Println("send error:", err)
			}
		}()
	}
}

func (f forwarder) start() {
	wg := &sync.WaitGroup{}
	wg.Add(f.workers)

	for i := 1; i <= f.workers; i++ {
		go func(i int) {
			f.work(i, wg)
		}(i)
	}

	wg.Wait()
}
