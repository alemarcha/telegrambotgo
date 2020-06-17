package main

import (
	"log"
	"time"
	"encoding/json"
	"fmt"
	tb "gopkg.in/tucnak/telebot.v2"
	"github.com/joho/godotenv"
	"os"

)

func main() {
	// load .env file
	godotenv.Load(".env")
	b, err := tb.NewBot(tb.Settings{
		// You can also set custom API URL.
		// If field is empty it equals to "https://api.telegram.org".
		Token:  os.Getenv("apikeybot"),
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/start", func(m *tb.Message) {
		jsonBody, _ := json.Marshal(m)

		fmt.Println(string(jsonBody))
		b.Send(m.Sender, "Hola, estamos integrando una plataforma de pago!")
	})

	b.Start()
}