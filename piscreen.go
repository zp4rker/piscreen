package main

import (
	"flag"
	"fmt"
	"image/color"
	"os/exec"
	"piscreen/keys"
	"piscreen/screens"
	"piscreen/spi"
	"piscreen/standby"
	"piscreen/util"
	"piscreen/vars"
	"time"
)

func main() {
	flag.IntVar(&vars.ListenDelay, "listen-delay", 300, "delay between button press listen")
	flag.BoolVar(&vars.Debug, "debug", false, "whether to print debug messages")
	flag.Parse()
	util.Debug(fmt.Sprintf("listen-delay = %vms\n", vars.ListenDelay))

	vars.Display = spi.NewST7789()
	defer vars.Display.Close()
	util.Debug("core initialised")

	vars.Running = true
	keys.Listen()
	standby.Observe()
	util.Debug("go routines initialised\n")

	vars.HomeScreen = screens.Home{}
	vars.MainMenu = screens.MainMenuInst()

	vars.CurrentScreen = vars.HomeScreen
	vars.PrevScreen = vars.CurrentScreen

	vars.Display.Clear(color.RGBA{})

	prevImage := vars.CurrentScreen.Render()

	for vars.Running {
		if vars.Display.IsAsleep() {
			continue
		}

		newImage := vars.CurrentScreen.Render()
		if util.ImageCmp(prevImage, newImage) {
			continue
		}

		prevImage = newImage
		vars.Display.ShowImage(newImage)
		util.Debug(fmt.Sprintf("\nnew image rendered @ %v", time.Now().Format("15:04:05")))
		time.Sleep(500 * time.Millisecond)
	}

	if vars.OnExit == "shutdown" {
		if err := exec.Command("sudo", "shutdown", "now").Run(); err != nil {
			panic(err)
		}
	} else if vars.OnExit == "restart app" {
		if err := exec.Command("sudo", "~zp4rker/lcd/go/piscreen").Start(); err != nil {
			panic(err)
		}
	} else if vars.OnExit == "reboot" {
		if err := exec.Command("sudo", "reboot").Run(); err != nil {
			panic(err)
		}
	} else {
		vars.Display.Clear(color.RGBA{})
	}
}
