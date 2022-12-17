package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	inputs := strings.Split(readLine(reader), " ")
	N, _ := strconv.ParseInt(inputs[0], 10, 64)
	K, _ := strconv.ParseInt(inputs[1], 10, 64)

	findAndPrint(N, K)
}

func findAndPrint(N, K int64) {
	printStr := ""
	for i := int64(125874); i <= N; i++ {
		var first []rune
		for j := int64(1); j <= K; j++ {
			cur := strconv.FormatInt(i*j, 10)

			if j == 1 {
				first = []rune(cur)
				sort.Slice(first, func(i, j int) bool {
					return first[i] < first[j]
				})
				printStr += cur + " "
			} else {
				if checkSame(first, []rune(cur)) {
					printStr += cur + " "

				} else {
					printStr = ""
					break
				}
			}
		}
		if printStr != "" {
			fmt.Println(printStr)
		}
	}
}

func checkSame(sorted []rune, unsorted []rune) bool {
	sort.Slice(unsorted, func(i, j int) bool {
		return unsorted[i] < unsorted[j]
	})
	return reflect.DeepEqual(sorted, unsorted)
}

func isIntegral(val float64) bool {
	return val == float64(int(val))
}
func isPentagon(n int64) bool {
	delta := 1 + 24*n
	n1 := (1 - math.Sqrt(float64(delta))) / 6
	n2 := (1 + math.Sqrt(float64(delta))) / 6

	return (n1 > 0 && isIntegral(n1)) || (n2 > 0 && isIntegral(n2))

}
func isTriangle(n int64) bool {
	delta := 1 + 8*n
	n1 := (-1 - math.Sqrt(float64(delta))) / 2
	n2 := (-1 + math.Sqrt(float64(delta))) / 2

	return (n1 > 0 && isIntegral(n1)) || (n2 > 0 && isIntegral(n2))

}

func isHexagonal(n int64) bool {
	delta := 1 + 8*n
	n1 := (1 - math.Sqrt(float64(delta))) / 4
	n2 := (1 + math.Sqrt(float64(delta))) / 4

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
