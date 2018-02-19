package model

import (
	"log"

	"github.com/line/line-bot-sdk-go/linebot"
)

type BotStruct struct {
	ChannelSecret string
	ChannelToken  string
}

func (b *BotStruct) Register() *linebot.Client {
	b.ChannelSecret = "06dfbfb76fef562140f6a5d949c45c6a"
	b.ChannelToken = "4sp/TOv5oBk0NUToWxnYznPn+fY+JnV5Am17sRYPvgymi40sfL8C/kFxxhN6nhUrkaJ+GbBlSvABMeMyYWZpB0mYjke2kFUQZVbyDZK4fiJOeM6yabiOTOCNt+GEzYaHDsDJ4uyMyaxoOvVuK0Ob2QdB04t89/1O/w1cDnyilFU="
	bot, err := linebot.New(b.ChannelSecret, b.ChannelToken)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return bot
}
