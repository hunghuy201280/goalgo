package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
)

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		inputText := scanner.Text()
		if inputText == "" {
			continue
		}
		input := strings.Split(strings.TrimSpace(inputText), ";")
		if len(input) != 3 {
			continue
		}
		result := convert(input[0], input[1], input[2])
		fmt.Println(result)
	}

	if scanner.Err() != nil {
		// Handle error.
	}
	//
	//for input := readLine(reader); input != ""; {
	//	input := strings.Split(strings.TrimSpace(readLine(reader)), ";")
	//	if len(input) != 3 {
	//		continue
	//	}
	//	result := convert(input[0], input[1], input[2])
	//	fmt.Println(result)
	//}

}

func convert(action string, typeStr string, input string) string {
	result := ""
	inputLen := len(input)
	//split
	if action == "S" {
		//remove ()
		input = strings.ReplaceAll(input, "()", "")
		for _, character := range input {
			if unicode.IsUpper(character) {
				result += " " + strings.ToLower(string(character))
			} else {
				result += string(character)
			}
		}
		return strings.TrimSpace(result)

	} else if action == "C" {
		nextUpper := false
		for index, character := range input {
			charStr := string(character)
			if index == 0 && typeStr == "C" {
				result += strings.ToUpper(charStr)
				continue
			}
			if charStr == " " {
				nextUpper = true
				continue
			}
			if nextUpper || (index == 0 && typeStr == "C") {
				result += strings.ToUpper(charStr)
				nextUpper = false
			} else {
				result += charStr
			}
			if index == inputLen-1 && typeStr == "M" {
				result += "()"
			}
		}
		return result
	}
	return ""
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}
