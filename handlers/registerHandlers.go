package handlers

import (
	"log"

	"github.com/line/line-bot-sdk-go/linebot"
)

func RegisterHandler(bot *linebot.Client, event *linebot.Event) error {
	res, err := bot.GetProfile(event.Source.UserID).Do()
	if err != nil {
		log.Fatal(err)
	}
	bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(res.DisplayName + " Anda telah terdaftar")).Do()
	return nil
}
