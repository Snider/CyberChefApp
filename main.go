package main

import (
	"embed"
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
	"github.com/wailsapp/wails/v2/pkg/runtime"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
)

//go:embed frontend/build/prod
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// BrowserOpenURL
	AppMenu := menu.NewMenu()
	FileMenu := AppMenu.AddSubmenu("File")
	FileMenu.AddText("&Open", keys.CmdOrCtrl("o"), func(_ *menu.CallbackData) {
		runtime.BrowserOpenURL(app.ctx, "http://localhost#recipe=To_Base64('A-Za-z0-9%2B/%3D')")
	})

	// Create application with options
	//  wails build -obfuscate
	err := wails.Run(&options.App{
		Title:            "CyberChef",
		Width:            1024,
		Height:           768,
		Assets:           assets,
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Menu:             AppMenu,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
