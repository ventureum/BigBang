package api

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

func BuildEndingPoint(baseURL URL, stage Stage, resource Resource) URL {
	endingPointURL := strings.Join([]string{string(baseURL), string(stage), string(resource)}, "/")
	return URL(endingPointURL)
}

func SendPost(message Message, url URL) *Message {
	bytesRepresentation, err := json.Marshal(message)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := http.Post(string(url), "application/json", bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		log.Fatalln(err)
	}

	var result Message

	json.NewDecoder(resp.Body).Decode(&result)

	return &result
}
