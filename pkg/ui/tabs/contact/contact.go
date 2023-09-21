package contact

import (
	"errors"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/dusnm/QRWiz/pkg/countries"
	dataContact "github.com/dusnm/QRWiz/pkg/data/contact"
	"github.com/dusnm/QRWiz/pkg/encoder"
	"github.com/dusnm/QRWiz/pkg/ui/components/qr"
	"github.com/dusnm/QRWiz/pkg/ui/core"
	"slices"
)

const (
	title = "Contact"
)

var selectedCountryCodeAlpha2 string
var selectedCountryCodeAlpha3 string

func Setup() *container.TabItem {
	name := widget.NewEntry()
	surname := widget.NewEntry()
	email := widget.NewEntry()
	phone := widget.NewEntry()
	website := widget.NewEntry()
	streetAndNumber := widget.NewEntry()
	postalCode := widget.NewEntry()
	state := widget.NewEntry()
	city := widget.NewEntry()
	country := countrySelector()
	button := widget.NewButton(
		"Encode",
		func() {
			if selectedCountryCodeAlpha2 == "" {
				dialog.NewError(errors.New("you must select a country"), core.Window).Show()
				return
			}

			c, err := dataContact.New(
				name.Text,
				surname.Text,
				postalCode.Text,
				streetAndNumber.Text,
				city.Text,
				state.Text,
				selectedCountryCodeAlpha2,
				selectedCountryCodeAlpha3,
				phone.Text,
				email.Text,
				website.Text,
			)

			if err != nil {
				dialog.NewError(err, core.Window).Show()
				return
			}

			b, err := encoder.Encode(c.String())
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
			container.NewBorder(
				nil,
				container.NewVBox(
					layout.NewSpacer(),
					container.NewCenter(button),
					layout.NewSpacer(),
				),
				nil,
				nil,
				container.NewGridWithColumns(
					2,
					container.NewVBox(
						widget.NewCard(
							"Personal",
							"Information pertaining to the person",
							container.NewVBox(
								container.NewGridWithColumns(
									2,
									container.NewVBox(
										widget.NewLabel("Name"),
										name,
									),
									container.NewVBox(
										widget.NewLabel("Surname"),
										surname,
									),
								),
								widget.NewLabel("Phone number"),
								phone,
								container.NewGridWithColumns(
									2,
									container.NewVBox(
										widget.NewLabel("Email"),
										email,
									),
									container.NewVBox(
										widget.NewLabel("Website"),
										website,
									),
								),
							),
						),
					),
					widget.NewCard(
						"Address",
						"Information pertaining to the person's place of residency",
						container.NewVBox(
							widget.NewLabel("Street and number"),
							streetAndNumber,
							container.NewGridWithColumns(
								2,
								container.NewVBox(
									widget.NewLabel("Postal/Zip code"),
									postalCode,
								),
								container.NewVBox(
									widget.NewLabel("State/Province"),
									state,
								),
							),
							container.NewGridWithColumns(
								2,
								container.NewVBox(
									widget.NewLabel("City"),
									city,
								),
								container.NewVBox(
									widget.NewLabel("Country"),
									country,
								),
							),
						),
					),
				),
			),
		),
	)
}

func countrySelector() *widget.Select {
	countryInfo, _ := countries.AllIndexedByName()
	options := make([]string, 0, len(countryInfo))

	for _, v := range countryInfo {
		options = append(options, v.Name)
	}

	slices.Sort(options)

	s := widget.NewSelect(options, func(s string) {
		selectedCountryCodeAlpha2 = countryInfo[s].CCA2
		selectedCountryCodeAlpha3 = countryInfo[s].CCA3
	})

	return s
}
