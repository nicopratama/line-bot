package handlers

import (
	"fmt"

	"github.com/line/line-bot-sdk-go/linebot"
)

func RegisterHandler(bot *linebot.Client, event *linebot.Event) error {
	fmt.Println(event.Source.UserID)
	bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("<html><head><body><title>Tes</title></body></head></html>"), linebot.NewTextMessage("Halo kakak")).Do()
	return nil
}
