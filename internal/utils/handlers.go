package utils

import "log"

func ErrorHandler(e error, panicMessage string) {
	if e == nil {
		return
	}

	if panicMessage == "" {
		log.Println(e)
		return
	} else {
		log.Printf("ERROR: %v\n", panicMessage)
	}

	log.Printf("ERROR! %v\n", e)
}
