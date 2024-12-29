package util

import (
	"fmt"
	"strings"
)

func TimeString(raw uint64) string {
	days := raw / 86400
	hours := raw % 86400 / 3600
	minutes := raw % 86400 % 3600 / 60
	seconds := raw % 86400 % 3600 % 60

	str := ""
	if days > 0 {
		str += fmt.Sprintf("%dd ", days)
	}
	if hours > 0 {
		str += fmt.Sprintf("%dh ", hours)
	}
	if minutes > 0 {
		str += fmt.Sprintf("%dm ", minutes)
	}
	if seconds > 0 {
		str += fmt.Sprintf("%ds", seconds)
	}

	return strings.TrimSpace(str)
}
