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
		n, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		res := largestPrimeFactor(n)
		fmt.Println(res)
	}
}

func largestPrimeFactor(n int64) int64 {
	temp := int64(2)
	maxFactor := int64(-1)
	var factors []int64
	for n > 1 {
		for n%temp == 0 {
			factors = append(factors, temp)
			if maxFactor < temp {
				maxFactor = temp
			}
			n /= temp
		}
		temp += 1

		if temp*temp > n && n != 1 {
			factors = append(factors, n)
			if maxFactor < n {
				maxFactor = n
			}
			break
		}
	}
	return maxFactor
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
