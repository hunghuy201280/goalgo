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

	tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	inputs := make([]int64, tTemp)
	for tItr := 0; tItr < int(tTemp); tItr++ {
		n, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		inputs[tItr] = n
	}

	MaxSize = 1001
	isPrimes := make([]bool, MaxSize+1)
	for i := range isPrimes {
		isPrimes[i] = true
	}
	primes := run(isPrimes)
	combinations := calc(MaxSize, primes)

	for _, input := range inputs {
		fmt.Println(combinations[input])
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

func calc(n int64, primes []int64) []int64 {
	combinations := make([]int64, n+1)
	combinations[0] = 1
	for _, prime := range primes {
		for i := int64(1); i < n+1; i++ {
			if i >= prime {
				combinations[i] += combinations[i-prime]
			}
		}
	}
	return combinations
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
	for i := int64(2); i < MaxSize; i++ {
		if primes[i] {
			result = append(result, i)
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
