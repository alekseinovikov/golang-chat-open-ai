package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"golang-chat-open-ai/core"
)

type uiService struct {
	apiKey      string
	chatService core.ChatService
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

	w.SetContent(
		container.New(
			layout.NewVBoxLayout(),
			u.form(),
		),
	)

	w.ShowAndRun()
}

func (u *uiService) form() *widget.Form {
	form := widget.NewForm()

	apiKeyEntry := widget.NewEntry()
	apiKeyEntry.SetPlaceHolder("Enter your api key here")

	welcomeMessageEntry := widget.NewMultiLineEntry()
	welcomeMessageEntry.SetPlaceHolder(u.chatService.GetDefaultWelcomeMessage())

	questionEntry := widget.NewMultiLineEntry()
	questionEntry.SetPlaceHolder("Please, explain me what the code does.")

	resultMessageEntry := widget.NewMultiLineEntry()
	resultMessageEntry.SetPlaceHolder("Result will be here.")

	form.Append("Api Key", apiKeyEntry)
	form.Append("Welcome Message", welcomeMessageEntry)
	form.Append("Question", questionEntry)
	form.Append("Result", resultMessageEntry)
	form.OnSubmit = func() {
		u.chatService.SetApiKey(apiKeyEntry.Text)
		result, err := u.chatService.Run(welcomeMessageEntry.Text, questionEntry.Text)
		if err != nil {
			resultMessageEntry.SetText(err.Error())
		} else {
			resultMessageEntry.SetText(result)
		}
	}

	form.SubmitText = "Ask Chat GPT"
	form.Refresh()

	return form
}
