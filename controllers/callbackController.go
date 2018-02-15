package controllers

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
)

func Callback(c *gin.Context) {
	bot, err := linebot.New("06dfbfb76fef562140f6a5d949c45c6a", "4sp/TOv5oBk0NUToWxnYznPn+fY+JnV5Am17sRYPvgymi40sfL8C/kFxxhN6nhUrkaJ+GbBlSvABMeMyYWZpB0mYjke2kFUQZVbyDZK4fiJOeM6yabiOTOCNt+GEzYaHDsDJ4uyMyaxoOvVuK0Ob2QdB04t89/1O/w1cDnyilFU=")
	if err != nil {
		log.Fatal(err)
	}

	events, err := bot.ParseRequest(c.Request)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			c.JSON(400, nil)
		} else {
			c.JSON(500, nil)
		}
		return
	}
	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				{
					bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do()
				}
			default:
				{

				}
			}
		}
	}
}
