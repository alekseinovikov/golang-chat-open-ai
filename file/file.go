package file

import (
	"golang-chat-open-ai/core"
	"io/ioutil"
	"strings"
)

type fileService struct {
}

func NewFileService() core.FileService {
	return &fileService{}
}

func (f *fileService) LoadAllTextFilesRecursivelyFromCurrentDirectory(fileExtensions []string) ([]core.FileContent, error) {
	files, err := ioutil.ReadDir("./")
	if err != nil {
		return nil, err
	}

	var textFiles []core.FileContent
	for _, file := range files {
		if file.IsDir() {
			continue
		}

		if hasAnySupportedSuffix(file.Name(), fileExtensions) {
			bytes, err := ioutil.ReadFile(file.Name())
			if err != nil {
				return nil, err
			}

			content := string(bytes)
			textFiles = append(textFiles, core.FileContent{
				Name:    file.Name(),
				Content: content,
			})
		}
	}

	return textFiles, nil
}

func hasAnySupportedSuffix(fileName string, supportedFileExtensions []string) bool {
	for _, ext := range supportedFileExtensions {
		if strings.HasSuffix(fileName, ext) {
			return true
		}
	}

	return false
}
