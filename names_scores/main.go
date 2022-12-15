package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	t := int32(tTemp)
	names := make([]string, tTemp)
	for tItr := 0; tItr < int(t); tItr++ {
		name := strings.TrimSpace(readLine(reader))
		names[tItr] = name
	}
	sort.Slice(names, func(i, j int) bool {
		return names[i] < names[j]
	})
	scores := calcScore(names)
	tTemp, err = strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	nQuery := int32(tTemp)
	for tItr := 0; tItr < int(nQuery); tItr++ {
		name := strings.TrimSpace(readLine(reader))
		fmt.Println(scores[name])
	}
}

func calcScore(names []string) map[string]int {
	scores := map[string]int{}
	for index, name := range names {
		scores[name] = calcScoreForName(name, index)
	}
	return scores
}
func calcScoreForName(name string, index int) int {
	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	score := 0
	for _, ru := range name {
		char := string(ru)
		score += strings.Index(alphabet, char) + 1
	}
	return score * (index + 1)
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
