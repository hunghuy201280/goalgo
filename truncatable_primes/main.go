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

var MaxSize int64 = 0

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	N := tTemp
	DigitCount := int64(math.Log10(float64(N)))
	MaxSize = int64(math.Pow(10, float64(DigitCount+1)))
	isPrimes := make([]bool, MaxSize)
	for i := range isPrimes {
		isPrimes[i] = true
	}
	run(isPrimes)
	calc(N, isPrimes)

}

func calc(n int64, isPrimes []bool) {
	sum := int64(0)

	for i := int64(11); i < n; i++ {
		if !isPrimes[i] {
			continue
		}
		isTruncatablePrime := IsTruncatablePrime([]rune(strconv.FormatInt(i, 10)), isPrimes)
		if isTruncatablePrime {
			sum += i
		}
	}
	fmt.Println(sum)
}
func removeAt(slice []rune, index int) []rune {
	return append(slice[:index], slice[index+1:]...)
}
func insert(a []rune, index int, value rune) []rune {
	if len(a) == index { // nil or empty slice or after last element
		return append(a, value)
	}
	a = append(a[:index+1], a[index:]...) // index < len(a)
	a[index] = value
	return a
}
func copySlice(slice []rune) []rune {
	cpy := make([]rune, len(slice))
	copy(cpy, slice)
	return cpy
}
func IsTruncatablePrime(a []rune, isPrimes []bool) bool {
	aCpyLeft := copySlice(a)
	aCpyRight := copySlice(a)

	for len(aCpyLeft) > 1 {
		aCpyLeft = removeAt(aCpyLeft, 0)
		aCpyRight = removeAt(aCpyRight, len(aCpyRight)-1)

		numValueLeft, _ := strconv.ParseInt(string(aCpyLeft), 10, 64)
		numValueRight, _ := strconv.ParseInt(string(aCpyRight), 10, 64)
		if !isPrimes[numValueLeft] || !isPrimes[numValueRight] {
			return false
		}
	}
	return true

}

func run(primes []bool) []int64 {
	var result []int64
	primes[0] = false
	primes[1] = false
	for i := 2; i*i < len(primes); i++ {
		if primes[i] {
			for j := i * i; j < len(primes); j += i {
				primes[j] = false
			}
		}
	}

	return result
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
