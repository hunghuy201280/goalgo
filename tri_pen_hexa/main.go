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
	a, _ := strconv.ParseInt(inputs[1], 10, 64)
	b, _ := strconv.ParseInt(inputs[2], 10, 64)

	findAndPrint(N, a, b)
}

func findAndPrint(n int64, a int64, b int64) {
	bNth := int64(0)
	for i := int64(1); i < n; {
		if a == 3 {

			bNth++

			if b == 5 {
				if isTriangle(i) {
					fmt.Println(i)
				}
				i += 3*bNth + 1
			} else {
				if bNth == 1 {
					bNth = 2
				}
				i = bNth*bNth + (bNth*bNth - bNth)
			}

		} else {
			if isPentagon(i) {
				fmt.Println(i)
			}
			bNth++
			if bNth == 1 {
				bNth = 2
			}
			i = bNth*bNth + (bNth*bNth - bNth)
		}
	}
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
