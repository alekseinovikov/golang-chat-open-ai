package main

import (
	"bufio"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"os"
	"strings"
)

func buildMessageBundle() []openai.ChatCompletionMessage {
	messages := make([]openai.ChatCompletionMessage, 0)
	messages = append(messages, createStartingMessage())
	messages = append(messages, createMessagesFrommFiles(loadAllTextFilesRecursivelyFromCurrentDirectory())...)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your question -> ")
	question, _ := reader.ReadString('\n')
	question = strings.Replace(question, "\n", "", -1)

	messages = append(messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: question,
	})

	return messages
}

func createStartingMessage() openai.ChatCompletionMessage {
	return openai.ChatCompletionMessage{
		Role: openai.ChatMessageRoleUser,
		Content: "Hello, I'm working on a project and I need your help. " +
			"I will send you all my files in the next messages. " +
			"Please, keep silent and let me finish. After that, please, answer my questions.",
	}
}

func createMessagesFrommFiles(files []fileContent) []openai.ChatCompletionMessage {
	messages := make([]openai.ChatCompletionMessage, 0)
	for _, f := range files {
		messages = append(messages, createMessageFromFile(f))
	}

	return messages
}

func createMessageFromFile(f fileContent) openai.ChatCompletionMessage {
	messageHeader := "//" + f.name + "\n\n"
	content := messageHeader + f.content
	return openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: content,
	}
}
