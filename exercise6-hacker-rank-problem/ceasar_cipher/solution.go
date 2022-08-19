package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'caesarCipher' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts following parameters:
 *  1. STRING s
 *  2. INTEGER k
 */

const (
	FLOOR_SMALL_LETTER    = 97
	CEILLING_SMALL_LETTER = 120
	FLOOR_BIG_LETTER      = 65
	CEILING_BIG_LETTER    = 90
)

func validateRune(r rune) (rune, bool) {
	if r >= FLOOR_BIG_LETTER && r <= 90 {
		return FLOOR_BIG_LETTER, true
	}
	if r >= FLOOR_SMALL_LETTER && r <= CEILLING_SMALL_LETTER {
		return FLOOR_SMALL_LETTER, true
	}
	return r, false
}

func countRotate(r rune, k int32) rune {
	if r+k > CEILING_BIG_LETTER && r+k < CEILLING_SMALL_LETTER {
		return FLOOR_BIG_LETTER - (CEILING_BIG_LETTER - r + 1)
	}
	return FLOOR_SMALL_LETTER - (CEILLING_SMALL_LETTER - (r - k))
}

func convertCipher(r rune, k int32) rune {
	if _, b := validateRune(r); !b {
		return 0
	}
	if c, b := validateRune(r + k); !b {
		return countRotate(c, k) // overflow has happened!!
	}
	return r + k
}

func caesarCipher(s string, k int32) string {
	// Write your code here
	var newCeasarString strings.Builder
	for _, c := range s {
		r := convertCipher(c, k)
		newCeasarString.Write([]byte(string(r)))
	}

	return newCeasarString.String()
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	// stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	// checkError(err)

	// defer stdout.Close()

	// writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

	// nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	// checkError(err)
	// n := int32(nTemp)

	s := readLine(reader)

	kTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	k := int32(kTemp)

	result := caesarCipher(s, int32(k))
	fmt.Println(result)

	// fmt.Fprintf(writer, "%s\n", result)

	// writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
