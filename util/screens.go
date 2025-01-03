package util

import (
	"fmt"
	"github.com/fogleman/gg"
	"piscreen/core"
	"piscreen/vars"
	"strings"
	"time"
)

func BaseScreen(footer bool) *gg.Context {
	context := gg.NewContext(240, 240)

	DarkGrey(context)
	context.DrawRectangle(0, 0, 240, 240)
	context.Fill()

	if err := context.LoadFontFace("JetBrainsMono.ttf", 17); err != nil {
		fmt.Printf("Failed to load font!")
	}

	if footer {
		White(context)
		context.DrawRectangle(0, 210, 240, 240)
		context.Fill()
		now := time.Now()
		timeFmt := "Mon 2 Jan 2006 15:04"
		if now.Unix()%2 == 0 {
			timeFmt = "Mon 2 Jan 2006 15 04"
		}
		DarkGrey(context)
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
		if strings.Contains(vars.CurrentScreen.Id(), "menu") {
			GoBackScreen()
		} else {
			vars.Display.ToggleSleep()
		}
		return true
	case "KEY3":
		if vars.CurrentScreen.Id() == "main_menu" {
			GoBackScreen()
		} else {
			OpenMainMenu()
		}
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
