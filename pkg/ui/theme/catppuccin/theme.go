package catppuccin

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"image/color"
	"sync"
)

var _ fyne.Theme = &Catppuccin{}

type (
	Catppuccin struct {
		mu     *sync.Mutex
		fonts  map[string]fyne.Resource
		colors map[fyne.ThemeVariant]map[fyne.ThemeColorName]color.Color
	}
)

func New() fyne.Theme {
	return &Catppuccin{
		mu:     &sync.Mutex{},
		fonts:  initFonts(),
		colors: initColors(),
	}
}

func (c *Catppuccin) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	c.mu.Lock()
	defer c.mu.Unlock()

	fallback := theme.DefaultTheme().Color(name, variant)
	colour, ok := c.colors[variant][name]
	if !ok {
		return fallback
	}

	if colour == (color.RGBA{}) {
		return fallback
	}

	return colour
}

func (c *Catppuccin) Font(style fyne.TextStyle) fyne.Resource {
	c.mu.Lock()
	defer c.mu.Unlock()

	if style.Monospace || style.Symbol || style.TabWidth != 0 {
		return theme.DefaultTheme().Font(style)
	}

	if style.Bold {
		return c.fonts[FontBold]
	}

	if style.Italic {
		return c.fonts[FontItalic]
	}

	return c.fonts[FontRegular]
}

func (c *Catppuccin) Icon(name fyne.ThemeIconName) fyne.Resource {
	c.mu.Lock()
	defer c.mu.Unlock()

	return theme.DefaultTheme().Icon(name)
}

func (c *Catppuccin) Size(name fyne.ThemeSizeName) float32 {
	c.mu.Lock()
	defer c.mu.Unlock()

	switch name {
	case theme.SizeNameText:
		return 18
	case theme.SizeNameHeadingText:
		return 28
	case theme.SizeNameSubHeadingText:
		return 22
	case theme.SizeNameCaptionText:
		return 15
	default:
		return theme.DefaultTheme().Size(name)
	}
}
