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
	t := tTemp
	inputs := make([]int64, t)
	for tItr := 0; tItr < int(t); tItr++ {
		nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		n := nTemp
		inputs[tItr] = n
	}

	maxDigit := findMax(inputs)
	cached := make([]int64, maxDigit+1)

	calcCache(cached, maxDigit)

	for _, value := range inputs {
		res := cached[value]
		fmt.Println(res)

	}
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

func calcCache(cached []int64, maxDigit int64) {
	phi := (1 + math.Sqrt(5)) / 2
	i := 1.0
	curDigit := int64(1)

	for cached[maxDigit] == 0 {
		digitCount := int64(i*math.Log10(phi) + math.Log10(1-math.Pow((1-phi)/phi, i)) - 0.5*math.Log10(5) + 1)
		if digitCount == curDigit {
			cached[curDigit] = int64(i)
			curDigit++
		}
		i++
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
