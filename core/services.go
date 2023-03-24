package core

type ChatService interface {
	SetApiKey(apiKey string)
	GetSupportedFileExtensions() []string
	GetDefaultWelcomeMessage() string
	LoadAndStoreFiles(path string, fileExtensions []string) error
	GetLoadedFileNames() []string
	Run(welcomeMessage string, question string) (string, error)
}

type FileContent struct {
	Name    string
	Content string
}

type FileService interface {
	LoadAllTextFilesRecursivelyFromCurrentDirectory(fileExtensions []string) ([]FileContent, error)
}

type UiService interface {
	Run()
}
