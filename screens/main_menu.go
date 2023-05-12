package screens

import (
	"image"
	"piscreen/util"
	"piscreen/vars"
)

type Menu struct {
	buttons []string
	focus   int
}

type MainMenu Menu

func MainMenuInst() *MainMenu {
	return &MainMenu{buttons: []string{
		"Home",
		"Something",
		"Something else",
		"Info",
		"Exit",
	}}
}

func (s *MainMenu) Id() string {
	return "main_menu"
}

func (s *MainMenu) Render() image.Image {
	context := util.BaseScreen(false)

	for i := 0; i < len(s.buttons); i++ {
		context.SetRGB(util.GGColor(0xFF, 0xFF, 0xFF))
		context.DrawRectangle(10, float64(10+i*46), 220, 36)
		if s.focus == i {
			context.Fill()
			context.SetRGB(util.GGColor(0x00, 0x00, 0x00))
		} else {
			context.Stroke()
		}

		context.DrawStringAnchored(s.buttons[i], 120, float64(10+i*46+18), 0.5, 0.5)
	}

	return context.Image()
}

func (s *MainMenu) Handle(key string) {
	if util.DefaultHandle(key) {
		s.focus = 0
		return
	}

	switch key {
	case "KEY_UP":
		if s.focus > 0 {
			s.focus--
		} else {
			s.focus = len(s.buttons) - 1
		}
	case "KEY_DOWN":
		if s.focus < len(s.buttons)-1 {
			s.focus++
		} else {
			s.focus = 0
		}
	case "KEY_PRESS":
		s.handleButton(s.buttons[s.focus])
		s.focus = 0
	}
}

func (s *MainMenu) handleButton(button string) {
	switch button {
	case "Home":
		util.GoHome()
		vars.PrevScreen = vars.CurrentScreen
	case "Info":
		util.ChangeScreen(Info{})
		vars.PrevScreen = vars.CurrentScreen
	}
}
