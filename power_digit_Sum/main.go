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
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		n := nTemp
		res := calc(n)
		fmt.Println(res)
	}
}

func calc(n int64) int64 {
	phi := (1 + math.Sqrt(5)) / 2
	i := 1.0
	if n > 3000 {
		i = 100000
	}

	for true {
		digitCount := int64(i*math.Log10(phi) + math.Log10(1-math.Pow((1-phi)/phi, i)) - 0.5*math.Log10(5) + 1)
		if digitCount == n {
			return int64(i)
		}
		i++
	}
	return 0
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
