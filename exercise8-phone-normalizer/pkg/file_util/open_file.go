package fileutil

import (
	"bufio"
	"os"
)

func ReadPhoneNumber(file *os.File) []string {
	scanner := bufio.NewScanner(file)
	var phone_numbers []string

	for scanner.Scan() {
		phone_numbers = append(phone_numbers, scanner.Text())
	}

	return phone_numbers
}
