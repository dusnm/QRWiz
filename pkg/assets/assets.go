package assets

import (
	"embed"
	"fyne.io/fyne/v2"
	"log"
)

//go:embed res/*
var assets embed.FS

func Icon() *fyne.StaticResource {
	i, err := assets.ReadFile("res/icon.png")
	if err != nil {
		log.Fatal(err)
	}

	return fyne.NewStaticResource("icon", i)
}

func ExampleQRCode() *fyne.StaticResource {
	i, err := assets.ReadFile("res/qr.png")
	if err != nil {
		log.Fatal(err)
	}

	return fyne.NewStaticResource("qr", i)
}
