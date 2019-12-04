package utils

import "log"

func Check(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}
