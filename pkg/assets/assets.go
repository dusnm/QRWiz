package assets

import (
	"embed"
	"log"

	"fyne.io/fyne/v2"
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

func ListOfCountries() *fyne.StaticResource {
	l, err := assets.ReadFile("res/countries.json")
	if err != nil {
		log.Fatal(err)
	}

	return fyne.NewStaticResource("countries", l)
}

func Font(style string) *fyne.StaticResource {
	f, err := assets.ReadFile("res/font-" + style + ".ttf")
	if err != nil {
		log.Fatal(err)
	}

	return fyne.NewStaticResource("font", f)
}
