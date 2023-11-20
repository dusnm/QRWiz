package catppuccin

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

func initColors() map[fyne.ThemeVariant]map[fyne.ThemeColorName]color.Color {
	// All color codes for all theme variants can be found at:
	// https://github.com/catppuccin/catppuccin
	//
	// The light variant uses Catppuccin latte
	// The dark variant uses Catppuccin mocha
	//
	// Many thanks to the authors of this wonderful theme ❤
	return map[fyne.ThemeVariant]map[fyne.ThemeColorName]color.Color{
		theme.VariantLight: {
			theme.ColorRed:                   color.RGBA{R: 0xd2, G: 0x0f, B: 0x39, A: 0xff}, // Red
			theme.ColorOrange:                color.RGBA{R: 0xfe, G: 0x64, B: 0x0b, A: 0xff}, // Peach
			theme.ColorYellow:                color.RGBA{R: 0xdf, G: 0x8e, B: 0x1d, A: 0xff}, // Yellow
			theme.ColorGreen:                 color.RGBA{R: 0x40, G: 0xa0, B: 0x2b, A: 0xff}, // Green
			theme.ColorBlue:                  color.RGBA{R: 0x1e, G: 0x66, B: 0xf5, A: 0xff}, // Blue
			theme.ColorPurple:                color.RGBA{R: 0x72, G: 0x87, B: 0xfd, A: 0xff}, // Lavender
			theme.ColorBrown:                 color.RGBA{R: 0xdc, G: 0x8a, B: 0x78, A: 0xff}, // Rosewater
			theme.ColorGray:                  color.RGBA{R: 0x9c, G: 0xa0, B: 0xb0, A: 0xff}, // Overlay0
			theme.ColorNameBackground:        color.RGBA{R: 0xef, G: 0xf1, B: 0xf5, A: 0xff}, // Base
			theme.ColorNameButton:            color.RGBA{R: 0xdf, G: 0x8e, B: 0x1d, A: 0xff}, // Yellow
			theme.ColorNameDisabledButton:    color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0x00},
			theme.ColorNameDisabled:          color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0x00},
			theme.ColorNameError:             color.RGBA{R: 0xd2, G: 0x0f, B: 0x39, A: 0xff}, // Red
			theme.ColorNameFocus:             color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0x00},
			theme.ColorNameForeground:        color.RGBA{R: 0x4c, G: 0x4f, B: 0x69, A: 0xff}, // Text
			theme.ColorNameHeaderBackground:  color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0x00},
			theme.ColorNameHover:             color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0x00},
			theme.ColorNameHyperlink:         color.RGBA{R: 0x1e, G: 0x66, B: 0xf5, A: 0xff}, // Blue
			theme.ColorNameInputBackground:   color.RGBA{R: 0xe6, G: 0xe9, B: 0xef, A: 0xff}, // Mantle
			theme.ColorNameInputBorder:       color.RGBA{R: 0xdf, G: 0x8e, B: 0x1d, A: 0xff}, // Yellow,
			theme.ColorNameMenuBackground:    color.RGBA{R: 0xef, G: 0xf1, B: 0xf5, A: 0xff}, // Base
			theme.ColorNameOverlayBackground: color.RGBA{R: 0x9c, G: 0xa0, B: 0xb0, A: 0xff}, // Overlay0
			theme.ColorNamePlaceHolder:       color.RGBA{R: 0x6c, G: 0x6f, B: 0x85, A: 0xff}, // Subtext0
			theme.ColorNamePressed:           color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0x00},
			theme.ColorNamePrimary:           color.RGBA{R: 0xfe, G: 0x64, B: 0x0b, A: 0xff}, // Peach
			theme.ColorNameScrollBar:         color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0x00},
			theme.ColorNameSelection:         color.RGBA{R: 0xdf, G: 0x8e, B: 0x1d, A: 0xff}, // Yellow
			theme.ColorNameSeparator:         color.RGBA{R: 0xdf, G: 0x8e, B: 0x1d, A: 0xff}, // Yellow
			theme.ColorNameShadow:            color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0x00},
			theme.ColorNameSuccess:           color.RGBA{R: 0x40, G: 0xa0, B: 0x2b, A: 0xff}, // Green
			theme.ColorNameWarning:           color.RGBA{R: 0xdf, G: 0x8e, B: 0x1d, A: 0xff}, // Yellow
		},
		theme.VariantDark: {
			theme.ColorRed:                   color.RGBA{R: 0xf3, G: 0x8b, B: 0xa8, A: 0xff}, // Red
			theme.ColorOrange:                color.RGBA{R: 0xfa, G: 0xb3, B: 0x87, A: 0xff}, // Peach
			theme.ColorYellow:                color.RGBA{R: 0xf9, G: 0xe2, B: 0xaf, A: 0xff}, // Yellow
			theme.ColorGreen:                 color.RGBA{R: 0xa6, G: 0xe3, B: 0xa1, A: 0xff}, // Green
			theme.ColorBlue:                  color.RGBA{R: 0x89, G: 0xb4, B: 0xfa, A: 0xff}, // Blue
			theme.ColorPurple:                color.RGBA{R: 0xcb, G: 0xa6, B: 0xf7, A: 0xff}, // Mauve
			theme.ColorBrown:                 color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0x00},
			theme.ColorGray:                  color.RGBA{R: 0x45, G: 0x47, B: 0x5a, A: 0xff}, // Surface1
			theme.ColorNameBackground:        color.RGBA{R: 0x1e, G: 0x1e, B: 0x2e, A: 0xff}, // Base
			theme.ColorNameButton:            color.RGBA{R: 0x11, G: 0x11, B: 0x1b, A: 0xff}, // Crust
			theme.ColorNameDisabledButton:    color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0x00},
			theme.ColorNameDisabled:          color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0x00},
			theme.ColorNameError:             color.RGBA{R: 0xf3, G: 0x8b, B: 0xa8, A: 0xff}, // Red
			theme.ColorNameFocus:             color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0x00},
			theme.ColorNameForeground:        color.RGBA{R: 0xcd, G: 0xd6, B: 0xf4, A: 0xff}, // Text
			theme.ColorNameHeaderBackground:  color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0x00},
			theme.ColorNameHover:             color.RGBA{R: 0x18, G: 0x18, B: 0x25, A: 0xff}, // Mantle
			theme.ColorNameHyperlink:         color.RGBA{R: 0x89, G: 0xb4, B: 0xfa, A: 0xff}, // Blue
			theme.ColorNameInputBackground:   color.RGBA{R: 0x18, G: 0x18, B: 0x25, A: 0xff}, // Mantle
			theme.ColorNameInputBorder:       color.RGBA{R: 0x11, G: 0x11, B: 0x1b, A: 0xff}, // Crust
			theme.ColorNameMenuBackground:    color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0x00},
			theme.ColorNameOverlayBackground: color.RGBA{R: 0x6c, G: 0x70, B: 0x86, A: 0xff}, // Overlay
			theme.ColorNamePlaceHolder:       color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0x00},
			theme.ColorNamePressed:           color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0x00},
			theme.ColorNamePrimary:           color.RGBA{R: 0xb4, G: 0xbe, B: 0xfe, A: 0xff}, // Lavender
			theme.ColorNameScrollBar:         color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0x00},
			theme.ColorNameSelection:         color.RGBA{R: 0x11, G: 0x11, B: 0x1b, A: 0xff}, // Crust
			theme.ColorNameSeparator:         color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0x00},
			theme.ColorNameShadow:            color.RGBA{R: 0x11, G: 0x11, B: 0x1b, A: 0xcc}, // Crust
			theme.ColorNameSuccess:           color.RGBA{R: 0xa6, G: 0xe3, B: 0xa1, A: 0xff}, // Green
			theme.ColorNameWarning:           color.RGBA{R: 0xf9, G: 0xe2, B: 0xaf, A: 0xff}, // Yellow
		},
	}
}
