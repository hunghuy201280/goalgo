package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
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

		//findWH(n)
	}

}

func findWH(target int64) {
	m := int64(1)
	for true {
		//delta:=1+16*float64(target)/float64(m*m+m)
		//s:=-1

		var n1 = (-1 + 4*math.Sqrt(float64(target)/float64(m*m+m))) / 2
		var n2 = (-1 - 4*math.Sqrt(float64(target)/float64(m*m+m))) / 2
		fmt.Println(n1, n2)
		m++
		time.Sleep(time.Duration(1000000))
	}
}

func checkCanPass(num int64, cached []bool) bool {
	cachedLen := len(cached)
	for i := 12; i < cachedLen; i++ {
		if cached[i] == false {
			continue
		}
		result := int64(i)

		for j := i; j < cachedLen; j++ {
			if cached[j] == false {
				continue
			}
			result += int64(j)

			if result == num {
				return true
			} else {
				result -= int64(j)
			}
		}
	}
	return false
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

func sumOfDivisor(num int64) bool {
	res := int64(1)
	for i := int64(2); i*2 <= num; i++ {
		if num%i == 0 {
			res += i
		}
	}
	if res > num {
		return true
	}
	return false

}

func run(cached []bool, maxInput int64) {
	for i := int64(12); i <= maxInput; i++ {
		cached[i] = sumOfDivisor(i)
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
