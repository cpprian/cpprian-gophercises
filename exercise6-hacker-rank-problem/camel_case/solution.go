package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

/*
 * Complete the 'camelcase' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts STRING s as parameter.
 */

func cutString(s string, r *regexp.Regexp) string {
	nextWordIndex := r.FindStringIndex(s)
	endNextWord := nextWordIndex[len(nextWordIndex)-1]
	if endNextWord == len(s) {
		return ""
	}

	return s[endNextWord:]
}

func camelcase(s string) int32 {
    // Write your code here
	stringSize := len(s)
	if stringSize < 1 {
		return 0
	}

	count := int32(1) // always start with at least a one word
	r, _ := regexp.Compile("[a-z]+")
	counter := cutString(s, r)
	if len(counter) == 0 {
		return count
	}

	r, _ = regexp.Compile("^[A-Z][a-z]+")
	for {
		counter = cutString(counter, r)
		count++
		if len(counter) == 0 {
			break
		}
	}

	return count
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    // stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    // checkError(err)

    // defer stdout.Close()

    // writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    s := readLine(reader)

    result := camelcase(s)
	fmt.Println(result)

    // fmt.Fprintf(writer, "%d\n", result)

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