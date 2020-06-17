package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"os"
	"github.com/joho/godotenv"
)

// webhook body
// https://core.telegram.org/bots/api#update
type webhookReqBody struct {
	Message struct {
		Text string `json:"text"`
		Chat struct {
			ID int64 `json:"id"`
			Title string `json:"title"`
			Type string `json:"type"`
			Username string `json:"username"`
		} `json:"chat"`
		From struct {
			ID int64 `json:"id"`
			Name string `json:"first_name"`
			Username string `json:"username"`
			IsBot bool `json:"is_bot"`
		} `json:"from"`
	} `json:"message"`
}


// Create a struct to conform to the JSON body of the send message request
// https://core.telegram.org/bots/api#sendmessage
type sendMessageReqBody struct {
	ChatID int64  `json:"chat_id"`
	Text   string `json:"text"`
}

// This handler is called everytime telegram sends us a webhook event
func Handler(res http.ResponseWriter, req *http.Request) {
	// load .env file
	godotenv.Load(".env")
	// First, decode the JSON response body
	body := &webhookReqBody{}
	if err := json.NewDecoder(req.Body).Decode(body); err != nil {
		fmt.Println("could not decode request body", err)
		return
	}

	// Check if the message contains the word "start"
	// else we do anythin
	jsonBody, _ := json.Marshal(body)
	fmt.Println(string(jsonBody))
	if !strings.Contains(strings.ToLower(body.Message.Text), "start") {
		return
	}

	// call start
	if err := start(body.Message.Chat.ID); err != nil {
		fmt.Println("error in sending reply:", err)
		return
	}

	// log a confirmation message if the message is sent successfully
	fmt.Println("reply sent")
}

//The below code deals with the process of sending a response message to the user
// start 
func start(chatID int64) error {
	// Create the request body struct
	reqBody := &sendMessageReqBody{
		ChatID: chatID,
		Text:   "Hola, estamos integrando una plataforma de pago!",
	}
	// Create the JSON body from the struct
	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		return err
	}
	// Send a post request with your token
	apikeybot := os.Getenv("apikeybot")
	urlSendMessage := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", apikeybot)

	res, err := http.Post(urlSendMessage, "application/json", bytes.NewBuffer(reqBytes))
	if err != nil {
		return err
	}
	
	if res.StatusCode != http.StatusOK {
		return errors.New("unexpected status" + res.Status)
	}

	return nil
}

// Start server
func main() {
	http.ListenAndServe(":3000", http.HandlerFunc(Handler))
}