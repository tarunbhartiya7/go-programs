/*
	Write a program to trim the string and remove all digits from a string and then reverse the string.

Example:
Input: "Hello 123 World 456"
Output: "dlroW olleH"
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func ModifyString(str string) string {
	str = strings.TrimSpace(str)
	re := regexp.MustCompile(`\d`)
	str = re.ReplaceAllString(str, "")
	result := ""
	for i := len(str) - 1; i >= 0; i-- {
		result += string(str[i])
	}
	return result
}

// func ModifyString(str string) string {
// 	str = strings.TrimSpace(str)
// 	result := ""
// 	for i := len(str) - 1; i >= 0; i-- {
// 		if str[i] >= '0' && str[i] <= '9' {
// 			continue
// 		}
// 		result += string(str[i])
// 	}
// 	return result
// }

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter a string: ")
	scanner.Scan()
	input := scanner.Text()

	result := ModifyString(input)
	fmt.Println("Result:", result)
}
