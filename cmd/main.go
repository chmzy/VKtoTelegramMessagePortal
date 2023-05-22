package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func main() {
	//botToken := "5810157518:AAEgr2V0GLfwb67AFsV4SLuqkkWLDJWbNSY"

	// Create Bot with debug on
	// Note: Please keep in mind that default logger may expose sensitive information, use in development only
	bot, err := telego.NewBot("token", telego.WithDefaultDebugLogger(), telego.WithHTTPClient(&http.Client{}))

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Get bot user
	botUser, err := bot.GetMe()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Bot user: %+v\n", botUser)

	// Get updates channel
	updates, _ := bot.UpdatesViaLongPolling(nil)

	// Stop reviving updates from update channel
	defer bot.StopLongPolling()

	// Loop through all updates when they came
	for update := range updates {
		// Check if update contains a message
		if update.Message != nil {

			chatID := tu.ID(update.Message.Chat.ID)

			if update.Message.Text == "create" {
				_, _ = bot.CreateForumTopic(&telego.CreateForumTopicParams{ChatID: chatID, Name: "test"})
			}

			_, _ = bot.SendMessage(&telego.SendMessageParams{ChatID: chatID, MessageThreadID: update.Message.MessageThreadID, Text: update.Message.Text})

			// Copy sent message back to the user
			//_, _ = bot.CopyMessage(
			//	tu.CopyMessage(
			//		chatID,
			//		chatID,
			//		update.Message.MessageID,
			//	),
			//)
		}
	}
}
