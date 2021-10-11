package parser

import (
	"encoding/json"
	"log"
)

// FormSubmit - struct for the form coming from the frontend
type FormSubmit struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Subject string `json:"subject"`
	Message string `json:"message"`
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
