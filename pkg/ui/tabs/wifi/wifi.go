package wifi

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	dataWifi "github.com/dusnm/QRWiz/pkg/data/wifi"
	"github.com/dusnm/QRWiz/pkg/encoder"
	"github.com/dusnm/QRWiz/pkg/ui/components/qr"
	"github.com/dusnm/QRWiz/pkg/ui/core"
)

const (
	title   = "Wi-Fi"
	secNone = "None"
	secWEP  = "WEP"
	secWPA  = "WPA/WPA2/WPA3"
)

func Setup() *container.TabItem {
	security := dataWifi.SEC_WPA
	hidden := false

	ssid := widget.NewEntry()
	password := widget.NewEntry()
	hiddenOptions := widget.NewCheck("Hidden", func(b bool) {
		hidden = b
	})

	securityOptions := widget.NewSelect(
		[]string{
			secNone,
			secWEP,
			secWPA,
		},
		func(s string) {
			switch s {
			case secNone:
				security = dataWifi.SEC_NONE
			case secWEP:
				security = dataWifi.SEC_WEP
			case secWPA:
				security = dataWifi.SEC_WPA
			}
		},
	)

	securityOptions.Selected = secWPA

	button := widget.NewButton(
		"Encode",
		func() {
			w, err := dataWifi.New(security, ssid.Text, password.Text, hidden)
			if err != nil {
				dialog.NewError(err, core.Window).Show()
				return
			}

			e, err := encoder.Encode(w.String())
			if err != nil {
				dialog.NewError(err, core.Window).Show()
				return
			}

			qr.Show(e)
		})

	return container.NewTabItem(
		title,
		container.NewPadded(
			container.NewVBox(
				widget.NewLabel("SSID"),
				ssid,
				widget.NewLabel("Password"),
				password,
				widget.NewLabel("Security"),
				securityOptions,
				hiddenOptions,
				container.NewCenter(button),
			),
		),
	)
}
