package core

type ChatService interface {
	SetApiKey(apiKey string)
	GetSupportedFileExtensions() []string
	SetSupportedFileExtensions(fileExtensions []string)
	GetDefaultWelcomeMessage() string
	LoadAndStoreFiles(path string, fileExtensions []string) error
	GetLoadedFileNames() []string
	Run(welcomeMessage string, question string) (string, error)
	GetSupportedModels() []string
	SetSelectedModel(model string)
}

type FileContent struct {
	Name    string
	Content string
}

type FileService interface {
	LoadAllTextFilesRecursivelyFromCurrentDirectory(path string, fileExtensions []string) ([]FileContent, error)
}

type UiService interface {
	Run()
}
