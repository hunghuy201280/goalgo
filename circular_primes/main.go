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

	for i := int64(2); i < n; i++ {
		if !isPrimes[i] {
			continue
		}
		isCircularPrime := IsCircularPrime([]rune(strconv.FormatInt(i, 10)), isPrimes)
		if isCircularPrime {
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
func IsCircularPrime(a []rune, isPrimes []bool) bool {
	NumberLength := len(a)
	checkCircularPrime := true
	aCpy := copySlice(a)

	for i := 0; i < NumberLength; i++ {

		aCpy = removeAt(aCpy, 0)
		aCpy = insert(aCpy, NumberLength-1, a[i])
		numValue, _ := strconv.ParseInt(string(aCpy), 10, 64)
		if !isPrimes[numValue] {
			checkCircularPrime = false
			break
		}

	}
	//if checkCircularPrime {
	//	fmt.Println("True", string(a))
	//
	//}
	return checkCircularPrime

}

func run(primes []bool) []int64 {
	var result []int64
	for i := 2; i*i < len(primes); i++ {
		if primes[i] {
			for j := i * i; j < len(primes); j += i {
				primes[j] = false
			}
		}
	}
	//for i := int64(2); i < MaxSize; i++ {
	//	if primes[i] {
	//		result = append(result, i)
	//	}
	//}

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
