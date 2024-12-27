package screens

import (
	"image"
	"piscreen/util"
)

type Home struct{}

func (s Home) Id() string {
	return "home"
}

func (s Home) Render() image.Image {
	context := util.BaseScreen(true)
	return context.Image()
}

func (s Home) Handle(key string) {
	if util.DefaultHandle(key) {
		return
	}
}
