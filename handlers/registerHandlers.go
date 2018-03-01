package handlers

import (
	"github.com/line/line-bot-sdk-go/linebot"
)

func RegisterHandler(bot *linebot.Client, replyToken string, source *linebot.EventSource) error {
	bot.ReplyMessage(replyToken, linebot.NewTextMessage("Anda telah terdaftar"), linebot.NewTextMessage("Halo kakak")).Do()
	return nil
}
