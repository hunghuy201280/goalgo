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

	tTemp, _ := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	findAndPrint(tTemp)
}

func findAndPrint(temp int64) {
	prefN := 3
	prefD := 2
	for i := int64(2); i <= temp; i++ {
		curN := prefN + 2*prefD
		curD := prefN + prefD
		prefN = curN
		prefD = curD
		curNStr := strconv.FormatInt(int64(curN), 10)
		curDStr := strconv.FormatInt(int64(curD), 10)
		if len(curNStr) > len(curDStr) {
			fmt.Println(i)
		}
	}
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
