package main

import (
	"io/ioutil"
	"log"
	"strings"
)

var supportedFileExtension = []string{".txt", ".md", ".gradle", ".java", ".kt", ".kts", ".ts", ".js", ".go", ".mod"}

type fileContent struct {
	name    string
	content string
}

func loadAllTextFilesRecursivelyFromCurrentDirectory() []fileContent {
	files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	var textFiles []fileContent
	for _, file := range files {
		if file.IsDir() {
			continue
		}

		if hasAnySupportedSuffix(file.Name()) {
			bytes, err := ioutil.ReadFile(file.Name())
			if err != nil {
				log.Println(err)
				continue
			}

			content := string(bytes)
			textFiles = append(textFiles, fileContent{
				name:    file.Name(),
				content: content,
			})
		}
	}

	return textFiles
}

func hasAnySupportedSuffix(fileName string) bool {
	for _, ext := range supportedFileExtension {
		if strings.HasSuffix(fileName, ext) {
			return true
		}
	}

	return false
}
