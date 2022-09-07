package main

import (
	"log"
	"os"

	dbutil "github.com/cpprian/cpprian-gophercises/exercise8-phone-normalizer/pkg/db_util"
	fileutil "github.com/cpprian/cpprian-gophercises/exercise8-phone-normalizer/pkg/file_util"
)

func main() {
	db_handle := dbutil.InitDbHandle()
	db_handle.OpenDB()
	defer db_handle.CloseDB()

	db_handle.PingDB()
	log.Println("Successfully connected!")

	f, err := os.Open("phones.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	for _, phone_number := range fileutil.ReadPhoneNumber(f) {
		db_handle.InsertPhoneNumber(phone_number)
	}

	db_handle.PrintPhoneNumbers()
}
