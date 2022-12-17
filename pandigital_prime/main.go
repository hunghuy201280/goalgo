package main

import (
	"bufio"
	"fmt"
	"io"
	"math/big"
	"os"
	"sort"
	"strconv"
	"strings"
)

var MaxSize int64 = 0

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	tTemp, _ := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	possibleNumber := getPossibleNumbers()
	sort.Slice(possibleNumber, func(i, j int) bool {
		return possibleNumber[i] < possibleNumber[j]
	})
	inputs := make([]int64, tTemp)
	for tItr := 0; tItr < int(tTemp); tItr++ {
		nTemp, _ := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		n := nTemp
		calc(n, possibleNumber)

		inputs[tItr] = n
	}

}

func findClosest(arr []int64, n int64, target int64) int64 {

	if target <= arr[0] {
		return arr[0]
	}
	if target >= arr[n-1] {
		return arr[n-1]
	}
	i := int64(0)
	j := n
	mid := int64(0)

	for i < j {
		mid = (i + j) / 2

		if arr[mid] == target {
			return arr[mid]
		}

		if target < arr[mid] {

			if mid > 0 &&
				target > arr[mid-1] {
				return getClosest(arr[mid-1], arr[mid], target, arr)
			}
			j = mid
		} else {
			if mid < n-1 &&
				target < arr[mid+1] {
				return getClosest(arr[mid], arr[mid+1], target, arr)
			}
			i = mid + 1
		}
	}
	return arr[mid]
}

func getClosest(val1 int64, val2 int64, target int64, arr []int64) int64 {

	if val1 < target && val2 < target {
		if target-val1 >= val2-target {
			return val2
		} else {
			return val1
		}
	} else if val1 > target {
		return val2
	} else {
		return val1
	}
}

// Perm calls f with each permutation of a.
func Perm(a []rune, f func([]rune)) {
	perm(a, f, 0)
}

// Permute the values at index i to len(a)-1.
func perm(a []rune, f func([]rune), i int) {
	if i > len(a) {
		f(a)
		return
	}
	perm(a, f, i+1)
	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		perm(a, f, i+1)
		a[i], a[j] = a[j], a[i]
	}
}
func getPossibleNumbers() []int64 {
	possiblePermutation := []string{"1", "12", "123", "1234", "12345", "123456", "1234567", "12345678", "123456789"}
	var res []int64
	for _, value := range possiblePermutation {

		//fmt.Println(value)

		Perm([]rune(value), func(temp []rune) {
			number, _ := strconv.ParseInt(string(temp), 10, 64)
			if isPrime(number) {
				res = append(res, number)
			}
			//res = append(res, number)

		})
	}
	return res
}
func calc(n int64, possibleNum []int64) {
	if n < possibleNum[0] {
		fmt.Println("-1")
		return
	}
	resIndex := sort.Search(len(possibleNum), func(i int) bool { return possibleNum[i] >= n })
	if resIndex < len(possibleNum) && possibleNum[resIndex] <= n {
		fmt.Println(possibleNum[resIndex])
	} else if possibleNum[resIndex-1] <= n {
		fmt.Println(possibleNum[resIndex-1])

	} else {
		fmt.Println("-1")

	}
	//
	//res := possibleNum[resIndex-1]
	//
	//findClosest(possibleNum, int64(len(possibleNum))-1, n)
	//fmt.Println(res)

}
func isPrime(n int64) bool {
	if big.NewInt(n).ProbablyPrime(0) {
		return true
	} else {
		return false
	}
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}
