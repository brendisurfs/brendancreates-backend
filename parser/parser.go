package parser

import (
	"encoding/json"
	"log"
)

type FormSubmit struct {
	Name    string
	Email   string
	Subject string
	Message string
}

// MessageParser - parses the received message from []bye to JSON.
func MessageParser(message []byte) FormSubmit {

	var msg FormSubmit

	err := json.Unmarshal(message, &msg)
	if err != nil {
		log.Fatal(err)
	}

	return msg
}
