package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var (
	headerHost     = "X-Sumo-Host"
	headerName     = "X-Sumo-Name"
	headerCategory = "X-Sumo-Category"
)

type sumopayload struct {
	host           string
	sourceName     string
	sourceCategory string
	fields         map[string]string
	lines          []string
}

func (p *sumopayload) body() io.Reader {
	return strings.NewReader(strings.Join(p.lines, "\n"))
}

func (p *sumopayload) send(client *http.Client, collector string) error {
	req, err := http.NewRequest("POST", collector, p.body())
	if err != nil {
		return err
	}

	if p.host != "" {
		req.Header.Add(headerHost, p.host)
	}
	if p.sourceName != "" {
		req.Header.Add(headerName, p.sourceName)
	}
	if p.sourceCategory != "" {
		req.Header.Add(headerCategory, p.sourceCategory)
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	log.Printf("sumo reply: %d %s\n", resp.StatusCode, reply)
	return nil
}
