package main

import (
	"golang-chat-open-ai/chat"
	"golang-chat-open-ai/file"
	"golang-chat-open-ai/ui"
	"os"
)

func main() {
	apiKey := os.Getenv("OPENAI_API_KEY")

	fileService := file.NewFileService()
	chatService := chat.NewChatService(fileService)

	ui.NewUiService(apiKey, chatService).Run()
}
