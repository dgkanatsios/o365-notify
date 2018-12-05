package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

const (
	EnvO365Webhook      = "O365_WEBHOOK"
	EnvO365Message      = "O365_MESSAGE"
	EnvO365AdaptiveCard = "O365_ADAPTIVECARD"
)

type Message struct {
	Text string `json:"text,omitempty"`
}

func main() {
	endpoint := os.Getenv(EnvO365Webhook)
	if endpoint == "" {
		fmt.Fprintln(os.Stderr, "URL is required")
		os.Exit(1)
	}
	text := os.Getenv(EnvO365Message)
	card := os.Getenv(EnvO365AdaptiveCard)
	if text == "" && card == "" {
		fmt.Fprintln(os.Stderr, "Message or Adaptive Card is required")
		os.Exit(1)
	}

	var encoded []byte
	if card == "" {
		msg := Message{
			Text: os.Getenv(EnvO365Message),
		}
		var err error
		encoded, err = json.Marshal(msg)
		if err != nil {
			fmt.Printf("Error in JSON conversion: %s", err.Error())
			os.Exit(1)
		}
	} else {
		encoded = []byte(card)
	}

	if err := send(endpoint, encoded); err != nil {
		fmt.Fprintf(os.Stderr, "Error sending message: %s\n", err)
		os.Exit(2)
	}
}

func send(endpoint string, encoded []byte) error {

	b := bytes.NewBuffer(encoded)
	res, err := http.Post(endpoint, "application/json", b)
	if err != nil {
		return err
	}

	if res.StatusCode >= 299 {
		return fmt.Errorf("error on message: %s - %#v", res.Status, res)
	}
	fmt.Println(res.Status)
	return nil
}
