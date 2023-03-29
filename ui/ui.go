package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"golang-chat-open-ai/core"
	"strings"
)

type uiService struct {
	apiKey      string
	chatService core.ChatService
	mainForm    *widget.Form
}

func NewUiService(apiKei string, chatService core.ChatService) core.UiService {
	return &uiService{
		apiKey:      apiKei,
		chatService: chatService,
	}
}

func (u *uiService) Run() {
	a := app.New()
	w := a.NewWindow("Open AI Code Chat")
	w.Resize(fyne.NewSize(800, 600))

	u.mainForm = u.submitForm()
	w.SetContent(
		container.New(
			layout.NewVBoxLayout(),
			u.findFilesForm(w),
			u.mainForm,
		),
	)

	w.ShowAndRun()
}

func (u *uiService) submitForm() *widget.Form {
	form := widget.NewForm()

	apiKeyEntry := widget.NewPasswordEntry()
	apiKeyEntry.SetPlaceHolder("Enter your api key here")

	modelSelection := widget.NewSelect(u.chatService.GetSupportedModels(), func(model string) {
		u.chatService.SetSelectedModel(model)
	})

	welcomeMessageEntry := widget.NewMultiLineEntry()
	welcomeMessageEntry.SetPlaceHolder(u.chatService.GetDefaultWelcomeMessage())
	welcomeMessageEntry.SetText(u.chatService.GetDefaultWelcomeMessage())

	questionEntry := widget.NewMultiLineEntry()
	questionEntry.SetPlaceHolder("Please, explain me what the code does.")
	questionEntry.SetText("Please, explain me what the code does.")

	resultMessageEntry := widget.NewMultiLineEntry()
	resultMessageEntry.SetPlaceHolder("Result will be here.")
	resultMessageEntry.SetText("Result will be here.")
	resultMessageEntry.SetMinRowsVisible(20)

	form.Append("Api Key", apiKeyEntry)
	form.Append("Model", modelSelection)
	form.Append("Welcome Message", welcomeMessageEntry)
	form.Append("Question", questionEntry)
	form.Append("Result", resultMessageEntry)
	form.OnSubmit = func() {
		resultMessageEntry.SetText("PROCESSING!!!! WAIT!!!!")
		form.SubmitText = "PROCESSING!!!! WAIT!!!!"

		u.chatService.SetApiKey(apiKeyEntry.Text)
		result, err := u.chatService.Run(welcomeMessageEntry.Text, questionEntry.Text)
		if err != nil {
			resultMessageEntry.SetText(err.Error())
		} else {
			resultMessageEntry.SetText(result)
		}

		form.SubmitText = "Ask Chat GPT"
	}

	form.SubmitText = "Ask Chat GPT"
	form.Disable()
	form.Refresh()

	return form
}

func (u *uiService) findFilesForm(parent fyne.Window) *widget.Form {
	form := widget.NewForm()

	filesPathEntry := widget.NewEntry()
	filesPathEntry.SetPlaceHolder("Path for files directory")
	filesPathEntry.SetText(".")

	form.AppendItem(widget.NewFormItem("Files Path", filesPathEntry))
	browseFilesDirectoryButton := widget.NewButton("Browse", func() {
		dialog.NewFolderOpen(func(uri fyne.ListableURI, err error) {
			if err != nil {
				dialog.ShowError(err, parent)
				return
			}
			if uri == nil {
				return
			}
			filesPathEntry.SetText(uri.Path())
		}, parent).Show()
	})
	form.Append("Browse", browseFilesDirectoryButton)

	fileExtensionsEntry := widget.NewEntry()
	fileExtensionsEntry.SetPlaceHolder(strings.Join(u.chatService.GetSupportedFileExtensions(), ","))
	fileExtensionsEntry.SetText(strings.Join(u.chatService.GetSupportedFileExtensions(), ","))
	form.Append("File Extensions For Scan (separate with comma)", fileExtensionsEntry)

	foundFiles := widget.NewMultiLineEntry()
	foundFiles.SetPlaceHolder("Found files will be here")
	form.Append("Found Files", foundFiles)

	form.OnSubmit = func() {
		splitted := strings.Split(fileExtensionsEntry.Text, ",")
		extentions := make([]string, 0)
		for _, ext := range splitted {
			extentions = append(extentions, strings.TrimSpace(ext))
		}

		u.chatService.SetSupportedFileExtensions(extentions)
		err := u.chatService.LoadAndStoreFiles(filesPathEntry.Text, extentions)
		if err != nil {
			foundFiles.SetText(err.Error())
		} else {
			loadedFileNames := u.chatService.GetLoadedFileNames()
			foundFiles.SetText(strings.Join(loadedFileNames, "\n"))
			if len(loadedFileNames) > 0 {
				u.mainForm.Enable()
			}
		}
	}
	form.SubmitText = "Find Files"

	form.Refresh()
	return form
}
