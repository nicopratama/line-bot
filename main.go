package main

import (
	"log"
	"net/http"
	"os"

	"github.com/line/line-bot-sdk-go/linebot"
)

func main() {
	// client := &http.Client{}
	bot, err := linebot.New("06dfbfb76fef562140f6a5d949c45c6a", "4sp/TOv5oBk0NUToWxnYznPn+fY+JnV5Am17sRYPvgymi40sfL8C/kFxxhN6nhUrkaJ+GbBlSvABMeMyYWZpB0mYjke2kFUQZVbyDZK4fiJOeM6yabiOTOCNt+GEzYaHDsDJ4uyMyaxoOvVuK0Ob2QdB04t89/1O/w1cDnyilFU=")
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/callback", func(w http.ResponseWriter, req *http.Request) {
		events, err := bot.ParseRequest(req)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				w.WriteHeader(400)
			} else {
				w.WriteHeader(500)
			}
			return
		}
		for _, event := range events {
			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do(); err != nil {
						log.Print(err)
					}
				}
			}
		}
	})

	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		log.Fatal(err)
	}
}
