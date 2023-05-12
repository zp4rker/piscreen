package vars

import (
	"piscreen/core"
	"time"
)

var Running = false

var Display core.Display
var LastActive = time.Now()
var CurrentScreen core.Screen
var PrevScreen core.Screen
var HomeScreen core.Screen
var MainMenu core.Screen

var Debug = false
var ListenDelay = 300
