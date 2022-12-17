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

	inputs := strings.Split(readLine(reader), " ")
	N, _ := strconv.ParseInt(inputs[0], 10, 64)
	K, _ := strconv.ParseInt(inputs[1], 10, 64)
	minP := int64(0)
	maxP := int64(1000000)
	Ps := calc(minP, maxP)
	findAndPrint(K, N, Ps)

}

func findAndPrint(K int64, N int64, ps []int64) {
	for i := K + 1; i < N; i++ {
		pMinus := ps[getPos(K+1, i)] - ps[getPos(K+1, i-K)]
		if isPentagon(pMinus) {
			fmt.Println(ps[getPos(K+1, i)])
			continue
		}
		pPlus := ps[getPos(K+1, i)] + ps[getPos(K+1, i-K)]
		if isPentagon(pPlus) {
			fmt.Println(ps[getPos(K+1, i)])
		}
	}
}

func getPos(minP, nth int64) int64 {
	return nth

}

func isIntegral(val float64) bool {
	return val == float64(int(val))
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

func isPentagon(n int64) bool {
	delta := 1 + 24*n
	n1 := (1 - math.Sqrt(float64(delta))) / 6
	n2 := (1 + math.Sqrt(float64(delta))) / 6

	return (n1 > 0 && isIntegral(n1)) || (n2 > 0 && isIntegral(n2))

}

func calc(minP int64, maxP int64) []int64 {
	n := maxP - minP + 1
	res := make([]int64, n)
	for i := int64(0); i < n; i++ {
		temp := i + minP
		res[i] = temp * (3*temp - 1) / 2
	}
	return res

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
