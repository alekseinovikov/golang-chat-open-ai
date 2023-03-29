package file

import (
	"golang-chat-open-ai/core"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type fileService struct {
}

func NewFileService() core.FileService {
	return &fileService{}
}

func (f *fileService) LoadAllTextFilesRecursivelyFromCurrentDirectory(path string, fileExtensions []string) ([]core.FileContent, error) {
	var textFiles []core.FileContent

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && hasAnySupportedSuffix(info.Name(), fileExtensions) {
			bytes, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}

			content := string(bytes)
			textFiles = append(textFiles, core.FileContent{
				Name:    info.Name(),
				Content: content,
			})
		}

		return nil
	})

	if err != nil {
		return nil, err
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
