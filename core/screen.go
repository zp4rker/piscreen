package core

import "image"

type Screen interface {
	Id() string
	Render() image.Image
	Handle(key string)
}
