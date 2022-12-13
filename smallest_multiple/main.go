package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		n := nTemp
		res := run(n)
		fmt.Println(res)
	}
}

func run(n int64) int64 {
	res := int64(1)
	for i := int64(1); i <= n; i++ {
		res = lcm(res, i)
	}
	return res
}
func lcm(a int64, b int64) int64 {
	return a * b / gcd(a, b)
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
func gcd(a int64, b int64) int64 {
	if a == 0 {

		return b
	}
	if b == 0 {

		return a
	}
	if a == b {
		return a
	}
	if a > b {
		return gcd(a-b, b)
	}
	return gcd(a, b-a)
}
