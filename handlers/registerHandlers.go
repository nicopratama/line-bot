package handlers

import (
	"github.com/line/line-bot-sdk-go/linebot"
)

func RegisterHandler(bot *linebot.Client, event *linebot.Event) error {
	bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(event.Source.UserID), linebot.NewTextMessage("Halo kakak")).Do()
	return nil
}
