package qr

import (
	"bytes"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/dusnm/QRWiz/pkg/ui/core"
	"image/png"
)

func Show(data []byte) {
	w := core.App.NewWindow("Preview")
	img := canvas.NewImageFromReader(bytes.NewReader(data), "qr")
	img.FillMode = canvas.ImageFillOriginal
	button := widget.NewButtonWithIcon(
		"Save",
		theme.DocumentSaveIcon(),
		func() {
			d := dialog.NewFileSave(
				func(closer fyne.URIWriteCloser, err error) {
					if err != nil {
						dialog.NewError(err, w).Show()
						return
					}

					if closer != nil {
						if err = png.Encode(closer, img.Image); err != nil {
							dialog.NewError(err, w).Show()
							return
						}
					}
				},
				w,
			)

			d.SetFileName("qr.png")
			d.Resize(fyne.Size{
				Width:  640,
				Height: 360,
			})

			d.Show()
		},
	)

	c := container.NewVBox(
		img,
		layout.NewSpacer(),
		container.NewCenter(button),
		layout.NewSpacer(),
	)

	w.SetContent(c)
	w.Show()
}
