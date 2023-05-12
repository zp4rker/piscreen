package screens

import (
	"image"
	"piscreen/util"
	"piscreen/vars"
)

type ExitMenu Menu

func ExitMenuInst() *ExitMenu {
	return &ExitMenu{buttons: []string{
		"Exit app",
		"Shutdown device",
		"Restart app",
		"Restart device",
		"Cancel",
	}}
}

func (s *ExitMenu) Id() string {
	return "exit_menu"
}

func (s *ExitMenu) Render() image.Image {
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

func (s *ExitMenu) Handle(key string) {
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

func (s *ExitMenu) handleButton(button string) {
	switch button {
	case "Exit app":
		vars.Running = false
	case "Shutdown device":
		vars.OnExit = "shutdown"
		vars.Running = false
	case "Restart app":
		vars.OnExit = "restart app"
		vars.Running = false
	case "Restart device":
		vars.OnExit = "reboot"
		vars.Running = false
	}
}
