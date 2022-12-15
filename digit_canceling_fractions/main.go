package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func miniMaxSum(arr []int32) {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	var minSum int64 = 0
	var maxSum int64 = 0

	for index, value := range arr {
		if index < 4 {
			minSum += int64(value)
		}
		if index > 0 {
			maxSum += int64(value)
		}
	}
	fmt.Printf("%d %d", minSum, maxSum)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")
	N, _ := strconv.ParseInt(arrTemp[0], 10, 64)
	K, _ := strconv.ParseInt(arrTemp[1], 10, 64)
	run(N, K)

}

func run(N int64, K int64) {
	start := int64(math.Pow(10, float64(N-1)) + 1)
	end := int64(math.Pow(10, float64(N)) - 1)
	startCancel := int64(math.Pow(10, float64(K-1)) + 1)
	endCancel := int64(math.Pow(10, float64(K)) - 1)
	numeratorSum := int64(0)
	denominatorSum := int64(0)
	for i := start; i <= end; i++ {
		numerator := i
		for j := numerator + 1; j <= end; j++ {
			denominator := j
			for cancelNum := startCancel; cancelNum <= endCancel; cancelNum++ {
				if cancelNum%10 == 0 {
					continue
				}
				if numerator%cancelNum == 0 && denominator%cancelNum == 0 {
					numeratorSum += numerator
					denominatorSum += denominator
				}
			}
		}
	}

	fmt.Printf("%d %d", numeratorSum, denominatorSum)

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
