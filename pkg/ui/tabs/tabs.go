package tabs

import (
	"fyne.io/fyne/v2/container"
	"github.com/dusnm/QRWiz/pkg/ui/tabs/contact"
	"github.com/dusnm/QRWiz/pkg/ui/tabs/email"
	"github.com/dusnm/QRWiz/pkg/ui/tabs/text"
	"github.com/dusnm/QRWiz/pkg/ui/tabs/wifi"
)

func Setup() *container.AppTabs {
	t := container.NewAppTabs(
		text.Setup(),
		wifi.Setup(),
		email.Setup(),
		contact.Setup(),
	)

	t.SetTabLocation(container.TabLocationTop)

	return t
}
