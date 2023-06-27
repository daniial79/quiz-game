package utils

import "log"

func LogErr(msg string, err error) {
	log.Fatalln(msg + ": " + err.Error())
}
