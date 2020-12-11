package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {

	var slice []int = make([]int, 3)
	var index int = 0
	for {

		fmt.Printf("enter a number !\n")
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {

			var inputStr string = scanner.Text()
			if inputStr == "X" {
				os.Exit(0)
			}
			input, _ := strconv.Atoi(inputStr)

			fmt.Printf("index :%d,  len(slice) : %d\n", index, len(slice))

			if index < len(slice) {
				for i := range slice {
					if slice[i] == 0 {
						slice[i] = input //add a number into slice
						break
					}
				}

			} else {
				slice = append(slice, input)
			}
			sort.Ints(slice) //sort
			fmt.Println(slice)
			index++
		}

	}
	prn
}
