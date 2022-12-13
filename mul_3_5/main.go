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
		n := int64(nTemp)
		res := calcSum(n)
		fmt.Println(res)
	}
}

func calcSum(n int64) int64 {
	N3 := n / 3
	N5 := n / 5
	N15 := n / 15
	if N3*3 == n {
		N3--
	}
	if N5*5 == n {
		N5--
	}
	if N15*15 == n {
		N15--
	}
	an3 := N3 * 3
	an5 := N5 * 5
	an15 := N15 * 15
	return ((3+an3)*N3 + (5+an5)*N5 - (15+an15)*N15) / 2
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
