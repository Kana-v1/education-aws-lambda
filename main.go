package main

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("dev.env")
	if err != nil {
		panic(err)
	}

	tgBotHost := "api.telegram.org"

	tgClient := NewClient(tgBotHost, os.Getenv("TELEGRAM_BOT_TOKEN"))

	chatID, _ := strconv.Atoi(os.Getenv("CHAT_ID"))

	tgClient.SendMessage(chatID, "message from bot")
}
