package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
)

type (
	payload struct {
		Time       string `json:"time"`
		Host       string `json:"host"`
		Source     string `json:"source"`
		SourceType string `json:"sourcetype"`
		Index      string `json:"index"`
		Event      struct {
			Line   string `json:"line"`
			Source string `json:"source"`
			Tag    string `json:"tag"`
		} `json:"event"`
	}
)

func parsePayloads(input []byte) []payload {
	decoder := json.NewDecoder(bytes.NewReader(input))
	result := []payload{}

	for {
		p := payload{}
		if err := decoder.Decode(&p); err != nil {
			if err == io.EOF {
				break
			}
			log.Println("decode error:", err)
			continue
		}
		result = append(result, p)
	}

	return result
}
