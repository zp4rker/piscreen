package keys

import "time"

type Key struct {
	Name           string
	Pin            string
	LastRegistered time.Time
}

func newKey(name, pin string) Key {
	return Key{name, pin, time.Now()}
}

var Keys = []Key{
	newKey("KEY_UP", "6"),
	newKey("KEY_DOWN", "19"),
	newKey("KEY_LEFT", "5"),
	newKey("KEY_RIGHT", "26"),
	newKey("KEY_PRESS", "13"),

	newKey("KEY1", "21"),
	newKey("KEY2", "20"),
	newKey("KEY3", "16"),
}
