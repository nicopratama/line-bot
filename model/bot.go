package model

import "github.com/line/line-bot-sdk-go/linebot"

type KitchenSink struct {
	bot         *linebot.Client
	appBaseURL  string
	downloadDir string
}
