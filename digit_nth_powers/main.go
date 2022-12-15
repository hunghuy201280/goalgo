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

	tTemp, _ := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	t := int32(tTemp)
	powArr := make([]int, 10)
	for i := 1; i <= 9; i++ {
		powArr[i] = int(math.Pow(float64(i), float64(t)))
	}

	res := 0
	for i := 100; i < 100000000; i++ {
		curNum := i
		sum := 0
		for curNum > 0 {
			temp := curNum % 10
			sum += powArr[temp]
			curNum /= 10
		}
		if i == sum {
			res += i
		}
	}
	fmt.Println(res)
}
func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}
