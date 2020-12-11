package main

import (
	"fmt"
)

func main() {
	for {

		fmt.Printf("put the float number : ")
		var input float32
		n, err := fmt.Scan(&input)
		if n <= 0 {
			fmt.Println("please input a number")
			fmt.Printf("err : %s", err)
		}

		fmt.Printf("converted to Integer : %d\n", int(input))
	}
}
