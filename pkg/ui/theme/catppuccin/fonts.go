package catppuccin

import (
	"fyne.io/fyne/v2"
	"github.com/dusnm/QRWiz/pkg/assets"
)

const (
	FontRegular = "regular"
	FontBold    = "bold"
	FontItalic  = "italic"
)

func initFonts() map[string]fyne.Resource {
	return map[string]fyne.Resource{
		FontRegular: assets.Font(FontRegular),
		FontBold:    assets.Font(FontBold),
		FontItalic:  assets.Font(FontItalic),
	}
}
