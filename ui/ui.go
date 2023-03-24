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

	w.SetContent(container.New(
		layout.NewVBoxLayout(),
		u.header(),
		u.body(),
		u.footer(),
	))

	w.ShowAndRun()
}

func (u *uiService) header() *fyne.Container {
	apiKeyLabel := widget.NewLabel("Api Key:")
	apiKeyEntry := widget.NewEntry()
	apiKeyEntry.SetPlaceHolder("Enter your api key here")
	apiKeyEntry.OnChanged = func(newValue string) {
		u.apiKey = newValue
		apiKeyLabel.Text = u.apiKey
		apiKeyLabel.Refresh()
	}

	return container.New(layout.NewHBoxLayout(),
		layout.NewSpacer(),
		container.New(layout.NewMaxLayout(), apiKeyLabel),
		container.New(layout.NewMaxLayout(), apiKeyEntry),
		layout.NewSpacer(),
	)
}

func (u *uiService) body() *fyne.Container {
	return container.NewHBox()
}

func (u *uiService) footer() *fyne.Container {
	return container.NewHBox()
}
