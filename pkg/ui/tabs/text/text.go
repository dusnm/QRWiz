package text

import (
	"errors"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/dusnm/QRWiz/pkg/encoder"
	"github.com/dusnm/QRWiz/pkg/ui/components/qr"
	"github.com/dusnm/QRWiz/pkg/ui/core"
)

const (
	title = "Text"
)

func Setup() *container.TabItem {
	input := widget.NewMultiLineEntry()
	button := widget.NewButton(
		"Encode",
		func() {
			if input.Text == "" {
				dialog.NewError(errors.New("text cannot be empty"), core.Window).Show()
				return
			}

			b, err := encoder.Encode(input.Text)
			if err != nil {
				dialog.NewError(err, core.Window).Show()
				return
			}

			qr.Show(b)
		},
	)
	return container.NewTabItem(
		title,
		container.NewPadded(
			container.NewVBox(
				widget.NewLabel("Text"),
				input,
				container.NewCenter(button),
			),
		),
	)
}
