package util

import (
	"fmt"
	"github.com/fogleman/gg"
	"piscreen/core"
	"piscreen/vars"
	"time"
)

func BaseScreen(footer bool) *gg.Context {
	context := gg.NewContext(240, 240)

	context.SetRGB(GGColor(0x35, 0x37, 0x39))
	context.DrawRectangle(0, 0, 240, 240)
	context.Fill()

	if err := context.LoadFontFace("JetBrainsMono.ttf", 17); err != nil {
		fmt.Printf("Failed to load font!")
	}

	if footer {
		context.SetRGB(GGColor(0xFF, 0xFF, 0xFF))
		context.DrawLine(0, 210, 240, 210)
		context.Stroke()
		now := time.Now()
		timeFmt := "Mon 2 Jan 2006 15:04"
		if now.Unix()%2 == 0 {
			timeFmt = "Mon 2 Jan 2006 15 04"
		}
		context.DrawStringAnchored(now.Format(timeFmt), 120, 225, 0.5, 0.5)
	}

	return context
}

func DefaultHandle(key string) bool {
	vars.LastActive = time.Now()

	if vars.Display.IsAsleep() {
		vars.Display.ToggleSleep()
		return true
	}

	switch key {
	case "KEY1":
		if vars.CurrentScreen == vars.PrevScreen {
			if vars.CurrentScreen.Id() == "home" {
				return false
			}
			GoHome()
		}
		GoBackScreen()
		return true
	case "KEY3":
		OpenMainMenu()
		return true
	}
	return false
}

func ChangeScreen(s core.Screen) {
	vars.PrevScreen = vars.CurrentScreen
	vars.CurrentScreen = s
}

func GoBackScreen() {
	vars.CurrentScreen = vars.PrevScreen
}

func GoHome() {
	ChangeScreen(vars.HomeScreen)
	vars.PrevScreen = vars.CurrentScreen
}

func OpenMainMenu() {
	ChangeScreen(vars.MainMenu)
}
