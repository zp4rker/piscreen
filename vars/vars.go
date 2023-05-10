package vars

import (
	"piscreen/core"
	"time"
)

var Running = false

var Display core.Display
var LastActive = time.Now()

var Debug = false
var ListenDelay = 300
