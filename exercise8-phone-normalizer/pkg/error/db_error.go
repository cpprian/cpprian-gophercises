package error_handler

import "log"

func DBError(err error) {
	if err != nil {
		log.Fatalf("Database error: %v\n", err)
	}
}
