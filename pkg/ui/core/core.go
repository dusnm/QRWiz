package core

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/dusnm/QRWiz/pkg/assets"
)

const (
	title = "QRWiz"
)

var App fyne.App
var Window fyne.Window

func SetupApp() {
	App = app.New()
	App.SetIcon(assets.Icon())
}

func SetupMainWindow(content fyne.CanvasObject) {
	Window = App.NewWindow(title)
	Window.Resize(fyne.Size{
		Width:  960,
		Height: 540,
	})

	Window.SetFixedSize(true)

	Window.SetContent(content)
}

func Run() {
	Window.ShowAndRun()
}