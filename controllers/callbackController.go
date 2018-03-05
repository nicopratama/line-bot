package controllers

import (
	"fmt"
	"log"

	"belajar/bot/handlers"
	"belajar/bot/services"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
)

func Callback(c *gin.Context) {
	firestoreClient := services.RegisterFirestore()
	// As an admin, the app has access to read and write all data, regradless of Security Rules
	ref := firestoreClient.NewRef("restricted_access/secret_document")
	var data map[string]interface{}
	if err := ref.Get(c, &data); err != nil {
		log.Fatalln("Error reading from database:", err)
	}
	fmt.Println(data)

	botModel := services.BotStruct{}
	bot := botModel.Register()
	events := handlers.CheckRequest(bot, c)
	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				{
					switch message.Text {
					case "/daftar":
						{
							if err := handlers.RegisterHandler(bot, event); err != nil {
								log.Print(err)
							}
						}
					case "/profile":
						{
							if err := HelpHandler(bot, event.ReplyToken, event.Source); err != nil {
								log.Print(err)
							}
						}
					default:
						{
							bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do()
							return
						}
					}
				}
			default:
				{
					return
				}
			}
		}
	}
}

func HelpHandler(bot *linebot.Client, replyToken string, source *linebot.EventSource) error {
	if _, err := bot.ReplyMessage(replyToken, linebot.NewTextMessage("Confirm Template")).Do(); err != nil {
		return err
	}
	return nil
}
