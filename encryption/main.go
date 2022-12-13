package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strings"
)

/*
 * Complete the 'encryption' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func encryption(s string) string {
	strCpy := s
	trimmed := strings.TrimSpace(s)
	L := int32(len(trimmed))
	rootL := math.Sqrt(float64(L))
	row := int32(math.Floor(rootL))
	col := int32(math.Ceil(rootL))

	for row*col < L {
		if row > col {
			col++
		} else {
			row++
		}
	}

	matrix := make([][]string, row)
	for i := range matrix {
		matrix[i] = make([]string, col)
	}
	curChar := ""
	for i := int32(0); i < row; i++ {
		for j := int32(0); j < col; j++ {
			if len(strCpy) > 1 {
				curChar, strCpy = string(strCpy[0]), strCpy[1:]
				matrix[i][j] = curChar
			} else if len(strCpy) == 1 {
				matrix[i][j] = strCpy
				break
			}
		}
	}
	res := ""
	for i := int32(0); i < col; i++ {
		for j := int32(0); j < row; j++ {
			res += matrix[j][i]
		}
		res += " "
	}
	return res
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	path, _ := os.Getwd()
	stdout, err := os.Create(path + "/encryption.txt")
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	s := readLine(reader)

	result := encryption(s)

	fmt.Fprintf(writer, "%s\n", result)

	writer.Flush()
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
