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
 * Complete the 'flippingBits' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts LONG_INTEGER n as parameter.
 */

func flippingBits(n int64) int64 {
	bitForm := fmt.Sprintf("%032b", n)
	converted := ""
	for i := 0; i < len(bitForm); i++ {
		char := string(bitForm[i])
		if char == "1" {
			converted += "0"
		} else {
			converted += "1"
		}
	}
	number, _ := strconv.ParseInt(converted, 2, 64)
	return number
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	path, err := os.Getwd()
	stdout, err := os.Create(path + "/temp.txt")
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	qTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	q := int32(qTemp)

	for qItr := 0; qItr < int(q); qItr++ {
		n, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)

		result := flippingBits(n)

		fmt.Fprintf(writer, "%d\n", result)
	}

	writer.Flush()
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
