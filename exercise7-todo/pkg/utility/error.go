package utility

import "log"

func IsErrorOccured(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
