package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	eliza "github.com/necrophonic/go-eliza"

	"github.com/gorilla/websocket"
)

// Message is a message event from the Slack websocket
type Message struct {
	ID      uint64 `json:"id"`
	UserID  string `json:"user"`
	Type    string `json:"type"`
	Channel string `json:"channel"`
	Text    string `json:"text"`
	SubType string `json:"subtype"`
}

var rtmDialer = websocket.Dialer{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	wsurl, id, err := startRTM()
	if err != nil {
		log.Fatalf("Failed to start Elizabot: %v\n", err)
	}

	ws, _, err := rtmDialer.Dial(wsurl, nil)
	if err != nil {
		log.Fatalf("Failed to start Elizabot: Websocket dial failed: %v\n", err)
	}

	log.Println("Started Elizabot [" + id + "]")

	for {
		message := Message{}
		if err = ws.ReadJSON(&message); err != nil {
			log.Printf("Error receiving message from Slack: %v\n", err)
			continue
		}

		if message.Type == "message" {
			// Skip if the message has our UserID. Stops the bot
			// responding to its own last message when starting up.
			if message.UserID == id {
				log.Println("Skipping my own message: " + message.Text)
				continue
			}

			message.Text, err = eliza.AnalyseString(message.Text)
			if err != nil {
				// Something went ary but we don't want Eliza to just have a breakdown,
				// so reply with a stock message and keep the bot active.
				log.Printf("Error analysing input [%s]: %v", message.Text, err)
				message.Text = "Can you rephrase that?"
			}
			if err = ws.WriteJSON(message); err != nil {
				// TODO Don't want to totally take out Eliza if something goes ary
				//      writing but need someway maybe to alert and halt if things
				//      really aren't working.
				log.Printf("Error attempting to write response so Slack: %v\n", err)
			}
		}
	}
}

type responseRtmStart struct {
	Ok    bool   `json:"ok"`
	Error string `json:"error"`
	URL   string `json:"url"`
	Self  struct {
		ID string `json:"ID"`
	}
}

func startRTM() (wsurl, id string, err error) {

	token := os.Getenv("ELIZA_BOT_TOKEN")
	if token == "" {
		return "", "", errors.New("Must set ELIZA_BOT_TOKEN environment variable")
	}

	resp, err := http.Get(fmt.Sprintf("https://slack.com/api/rtm.start?token=%s", token))
	defer resp.Body.Close()
	if err != nil {
		return "", "", err
	}
	if resp.StatusCode != 200 {
		return "", "", fmt.Errorf("API request failed with code [%d]", resp.StatusCode)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", "", err
	}

	var responseObject responseRtmStart
	if err = json.Unmarshal(body, &responseObject); err != nil {
		return "", "", err
	}

	if !responseObject.Ok {
		return "", "", fmt.Errorf("Slack error: %s", responseObject.Error)
	}

	return responseObject.URL, responseObject.Self.ID, nil
}
