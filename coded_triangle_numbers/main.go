package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	for tItr := 0; tItr < int(tTemp); tItr++ {
		n, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		fmt.Println(calc(n))
	}

}
func isIntegral(val float64) bool {
	return val == float64(int(val))
}
func findMax(arr []int64) int64 {
	max := int64(-1)
	for _, value := range arr {
		if max < value {
			max = value
		}
	}
	return max
}

func calc(n int64) int64 {
	delta := 1 + 8*n
	n1 := (-1 + math.Sqrt(float64(delta))) / 2
	n2 := (-1 - math.Sqrt(float64(delta))) / 2

	if n1 > 0 && isIntegral(n1) {
		return int64(n1)
	}
	if n2 > 0 && isIntegral(n2) {
		return int64(n2)
	}
	return -1

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
