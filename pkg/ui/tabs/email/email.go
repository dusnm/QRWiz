package email

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/dusnm/QRWiz/pkg/encoder"
	"github.com/dusnm/QRWiz/pkg/ui/components/qr"
	"github.com/dusnm/QRWiz/pkg/ui/core"

	dataEmail "github.com/dusnm/QRWiz/pkg/data/email"
)

const (
	title = "E-Mail"
)

func Setup() *container.TabItem {
	to := widget.NewEntry()
	cc := widget.NewEntry()
	bcc := widget.NewEntry()
	subject := widget.NewEntry()
	body := widget.NewMultiLineEntry()
	spacer := layout.NewSpacer()
	button := widget.NewButton(
		"Encode",
		func() {
			m, err := dataEmail.New(
				to.Text,
				cc.Text,
				bcc.Text,
				subject.Text,
				body.Text,
			)

			if err != nil {
				dialog.NewError(err, core.Window).Show()
				return
			}

			b, err := encoder.Encode(m.String())
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
				widget.NewLabel("To"),
				to,
				widget.NewLabel("CC (comma separated)"),
				cc,
				widget.NewLabel("BCC (comma separated)"),
				bcc,
				widget.NewLabel("Subject"),
				subject,
				widget.NewLabel("Body"),
				body,
				spacer,
				container.NewCenter(button),
				spacer,
			),
		),
	)
}
