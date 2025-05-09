package utils

import "time"

func MAKEID(name string) string {
	ts := time.Now().UTC().Format("20060102150405")
	return ts + "_" + name
}
