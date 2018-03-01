package handlers

import (
	"fmt"

	"github.com/line/line-bot-sdk-go/linebot"
)

func RegisterHandler(bot *linebot.Client, replyToken string, source *linebot.EventSource, event linebot.Event) error {
	fmt.Println(event.Source)
	bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("Anda telah terdaftar"), linebot.NewTextMessage("Halo kakak")).Do()
	return nil
}
