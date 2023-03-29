package chat

import (
	"context"
	"github.com/sashabaranov/go-openai"
	"golang-chat-open-ai/core"
	"strings"
)

var defaultSupportedFileExtension = []string{".txt", ".md", ".gradle", ".java", ".kt", ".kts", ".ts", ".js", ".go", ".mod", ".rs"}

type chatService struct {
	apiKey                 string
	fileService            core.FileService
	loadedFiles            []core.FileContent
	supportedFileExtension []string
	selectedModel          string
}

func NewChatService(fileService core.FileService) core.ChatService {
	return &chatService{
		fileService:            fileService,
		supportedFileExtension: defaultSupportedFileExtension,
	}
}

func (c *chatService) GetSupportedFileExtensions() []string {
	return c.supportedFileExtension
}

func (c *chatService) SetSupportedFileExtensions(fileExtensions []string) {
	c.supportedFileExtension = fileExtensions
}

func (c *chatService) SetApiKey(apiKey string) {
	c.apiKey = apiKey
}

func (c *chatService) LoadAndStoreFiles(path string, fileExtensions []string) error {
	files, err := c.fileService.LoadAllTextFilesRecursivelyFromCurrentDirectory(path, fileExtensions)
	if err != nil {
		return err
	}

	c.loadedFiles = files
	return nil
}

func (c *chatService) GetLoadedFileNames() []string {
	var fileNames []string
	for _, file := range c.loadedFiles {
		fileNames = append(fileNames, file.Name)
	}

	return fileNames
}

func (c *chatService) GetDefaultWelcomeMessage() string {
	return "Hello, I'm working on a project and I need your help. \n" +
		"I will send you all my files in the next messages. \n" +
		"Please, keep silent and let me finish. After that, please, answer my questions."
}

func (c *chatService) GetSupportedModels() []string {
	return []string{openai.GPT4, openai.GPT3Dot5Turbo}
}

func (c *chatService) SetSelectedModel(model string) {
	c.selectedModel = model
}

func (c *chatService) Run(welcomeMessage string, question string) (string, error) {
	messages := make([]openai.ChatCompletionMessage, 0)
	messages = append(messages, createUserMessage(welcomeMessage))
	messages = append(messages, c.createMessagesFrommFiles()...)
	messages = append(messages, createUserMessage(strings.Replace(question, "\n", "", -1)))

	client := openai.NewClient(c.apiKey)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:    c.selectedModel,
			Messages: messages,
		},
	)

	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}

func (c *chatService) createMessagesFrommFiles() []openai.ChatCompletionMessage {
	messages := make([]openai.ChatCompletionMessage, 0)
	for _, f := range c.loadedFiles {
		messages = append(messages, createMessageFromFile(f))
	}

	return messages
}

func createMessageFromFile(f core.FileContent) openai.ChatCompletionMessage {
	messageHeader := "//" + f.Name + "\n\n"
	content := messageHeader + f.Content
	return createUserMessage(content)
}

func createUserMessage(message string) openai.ChatCompletionMessage {
	return openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: message,
	}
}
