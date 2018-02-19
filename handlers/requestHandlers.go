package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
)

func CheckRequest(bot *linebot.Client, ctx *gin.Context) []*linebot.Event {
	events, err := bot.ParseRequest(ctx.Request)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			ctx.JSON(400, nil)
		} else {
			ctx.JSON(500, nil)
		}
		return nil
	}
	return events
}
