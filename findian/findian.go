package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Printf("Enter String\n")

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		var searchString string = scanner.Text()
		lowerCase := strings.ToLower(searchString)
		if isFoundain(lowerCase) {
			fmt.Printf("Found ! \n")
		} else {
			fmt.Printf("Not Found ! \n")
		}
	}
}
func isFoundain(str string) bool {

	if strings.Index(str, "i") == 0 && strings.Index(str, "n") == len(str)-1 && strings.IndexAny(str, "a") != -1 {
		return true
	} else {
		return false
	}
}
