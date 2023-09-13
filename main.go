package main

import (
	"github.com/dusnm/QRWiz/pkg/ui/core"
	"github.com/dusnm/QRWiz/pkg/ui/tabs"
)

func main() {
	core.SetupApp()
	core.SetupMainWindow(tabs.Setup())
	core.Run()
}
