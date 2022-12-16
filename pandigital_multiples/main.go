package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var MaxSize int64 = 0

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	tTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")
	N, _ := strconv.ParseInt(tTemp[0], 10, 64)
	K, _ := strconv.ParseInt(tTemp[1], 10, 64)

	cached := make([]bool, N+1)
	calcCached(cached, K)

	for i := int64(1); i < N; i++ {
		if cached[i] {
			fmt.Println(i)
		}
	}
}

func isPan(num int64, k int) bool {
	i := int64(1)
	mul := ""
	for true {
		mul += strconv.FormatInt(num*i, 10)
		i++

		if len(mul) >= k {
			break
		}
	}

	if len(mul) > k {
		return false
	}

	for i := 1; i <= k; i++ {
		iRune := strconv.Itoa(i)
		if !strings.Contains(mul, iRune) {
			return false
		}
	}
	return true

}

func calcCached(cached []bool, k int64) {
	maxLen := len(cached)
	cached[0] = false
	cached[1] = false

	for i := 2; i < maxLen; i++ {
		cached[i] = isPan(int64(i), int(k))
	}
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}
