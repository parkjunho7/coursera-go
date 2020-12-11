package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type stName struct {
	firstName string
	lastName  string
}

func main() {
	const STRSIZE int = 20
	const INITSLSIZE int = 5
	fmt.Printf("Enter File name : ")
	fileScanner := bufio.NewScanner(os.Stdin)
	fileScanner.Scan()
	file, err := os.Open(fileScanner.Text())
	if err == nil {
		scanner := bufio.NewScanner(file)
		var slice = make([]stName, 5)
		var index int
		for {
			if scanner.Scan() {
				nameText := scanner.Text()
				split := strings.Split(nameText, " ")

				var newName stName
				newName.firstName = stringLimit(split[0], STRSIZE)
				newName.lastName = stringLimit(split[1], STRSIZE)
				if index < INITSLSIZE {
					slice[index] = newName
				} else {
					slice = append(slice, newName)
				}
				index++
			} else {
				break
			}
		}
		for i := 0; i < len(slice); i++ {
			fmt.Printf("[%d]First Name : %s, Last Name : %s\n", i, slice[i].firstName, slice[i].lastName)
		}
	}
}

func stringLimit(str string, size int) string {

	reader := strings.NewReader(str)
	buf := make([]byte, size)
	n, _ := io.ReadAtLeast(reader, buf, size)
	if n != 0 {
		return string(buf)
	}
	return str

}
