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
 * Complete the 'countingValleys' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER steps
 *  2. STRING path
 */

func countingValleys(steps int32, path string) int32 {
	// Write your code here
	curHeight := int32(0)
	down := false
	res := 0
	for _, value := range path {
		char := string(value)
		if char == "D" {
			curHeight--
		} else {
			curHeight++
		}

		//down hill
		if curHeight == -1 {
			down = true
		}
		if curHeight == 0 && down {
			res++
			down = false
		}

	}
	return int32(res)
}

func pangrams(s string) string {
	// Write your code here
	isPan := true
	alphabets := "abcdefghijklmnopqrstuvwxyz"
	s = strings.ToLower(s)
	for _, value := range alphabets {
		char := string(value)
		if !strings.Contains(s, char) {
			isPan = false
			break
		}
	}
	if isPan {
		return "pangram"
	}
	return "not pangram"

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	stepsTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	steps := int32(stepsTemp)

	path := readLine(reader)

	result := countingValleys(steps, path)

	fmt.Fprintf(writer, "%d\n", result)

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
